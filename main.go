package main

import "jobschedulerapi/api"

// @title Job Scheduler API
// @version 1.0
// @description This is a job scheduler API for scheduling jobs and running them at a specified time

// @host localhost:8080
// @BasePath /api/v1
func main() {
	api.Run()
}
