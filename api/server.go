package api

import (
	"fmt"
	"jobschedulerapi/api/app"
	"jobschedulerapi/api/scheduler"
	_ "jobschedulerapi/docs"
	"jobschedulerapi/util/migration"
	"log"
)

func Run() {
	migration.Migrate()

	stopScheduler := make(chan bool)
	go scheduler.CheckAndRunJobs(stopScheduler)

	app := app.NewApp()
	port := 8080
	fmt.Printf("Server is running at http://localhost:%d\n", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))

}
