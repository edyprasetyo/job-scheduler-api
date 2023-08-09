package exception

import (
	"jobschedulerapi/api/handler"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if err != nil {
		errorMessage := err.Error()
		return handler.InternalError(ctx, errorMessage)
	}
	return nil
}
