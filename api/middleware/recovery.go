package middleware

import (
	"fmt"
	"jobschedulerapi/api/handler"
	"log"

	"github.com/gofiber/fiber/v2"
)

func RecoveryMiddleware(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic: %v\n", r)

			if err, ok := r.(error); ok {
				errorMessage := err.Error()

				fmt.Printf("Error: %+v\n", err)

				fmt.Printf("Panic occurred: %s\n", errorMessage)

				handler.InternalError(c, errorMessage)
			} else {
				fmt.Printf("Panic occurred: %v\n", r)

				handler.InternalError(c, fmt.Sprintf("%v", r))
			}
		}
	}()
	return c.Next()
}
