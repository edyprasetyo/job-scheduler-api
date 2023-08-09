package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ErrorMiddleware(c *fiber.Ctx) error {
	if err := c.Next(); err != nil {
		// Access the error message from the context
		errorMessage := err.Error()
		fmt.Printf("Error occurred: %s\n", errorMessage)

		// You can handle the error or customize the response here
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errorMessage,
		})
	}

	return nil
}
