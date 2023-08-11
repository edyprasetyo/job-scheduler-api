package scheduler

import (
	"jobschedulerapi/application/repository_impl"
	"jobschedulerapi/application/service"
	"jobschedulerapi/domain/model"
	"jobschedulerapi/domain/repository"
	"jobschedulerapi/infrastructure/service_impl"
	"jobschedulerapi/util/datetime"
	"log"
	"net/http"
	"time"
)

var intervalInSeconds = 20

func CheckAndRunJobs(stop chan bool) {
	db := service_impl.NewDatabase()
	CheckJob(db, stop)
	ticker := time.NewTicker(time.Duration(intervalInSeconds) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			CheckJob(db, stop)
		case <-stop:
			log.Println("Stopping job scheduler")
			return
		}
	}
}

func CheckJob(db service.Database, stop chan bool) {
	log.Println("Checking jobs")
	jobsRepository := repository_impl.NewJobsRepository(db)
	jobs, err := jobsRepository.FetchAll("IsExecuted=0", nil)
	if err != nil {
		log.Println(err)
		return
	}
	for _, job := range jobs {
		RunJob(job, jobsRepository, stop)
	}
}

func ExecuteJob(job model.TrJobsMdl, repository repository.JobsRepository, stop chan bool) {
	res, err := http.Get(job.APIUrl)
	if err != nil {
		log.Println(err)
		job.IsExecuted = true
		job.IsSuccess = false
		err = repository.Update(&job)
		if err != nil {
			log.Println(err)
		}
		return
	}
	defer res.Body.Close()
	job.IsExecuted = true
	job.IsSuccess = true
	err = repository.Update(&job)
	if err != nil {
		log.Println(err)
		return
	}
}

func RunJob(job model.TrJobsMdl, repository repository.JobsRepository, stop chan bool) {
	now := datetime.Now()
	isNeedTobeExecuted := datetime.Before(job.ExecutedAt, now)
	if isNeedTobeExecuted {
		log.Println("Job will be executed now")
		ExecuteJob(job, repository, stop)
	}

	diffInSeconds := datetime.DiffInSeconds(job.ExecutedAt, now)
	if diffInSeconds < intervalInSeconds && diffInSeconds > 0 {
		log.Println("Job will be executed in", diffInSeconds, "seconds")
		RunInQueue(job, repository, diffInSeconds, stop)
	}
}

func RunInQueue(job model.TrJobsMdl, repository repository.JobsRepository, delayInSeconds int, stop chan bool) {
	delayDuration := time.Duration(delayInSeconds) * time.Second
	timerChan := time.After(delayDuration)

	for {
		select {
		case <-timerChan:
			ExecuteJob(job, repository, stop)
			log.Println("Job executed from queue")
			return
		case <-stop:
			log.Println("Stopping job scheduler")
			return
		}
	}
}
