package router

import (
	"fmt"
)

func Run() {
	router := SetupRouter()
	port := 8080
	fmt.Printf("Server is running at http://localhost:%d\n", port)
	router.Run(fmt.Sprintf(":%d", port))
}
