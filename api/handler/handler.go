package handler

import (
	"jobschedulerapi/domain/extmodel"

	"github.com/gofiber/fiber/v2"
)

func Success(c *fiber.Ctx, obj interface{}) error {
	res := extmodel.ApiResponse{
		StatusCode: 200,
		Success:    true,
		Result:     obj,
	}
	return c.Status(200).JSON(res)
}

func ValidationError(c *fiber.Ctx, obj interface{}) error {
	res := extmodel.ApiResponse{
		StatusCode: 400,
		Success:    false,
		Result:     obj,
	}
	return c.Status(400).JSON(res)
}

func InternalError(c *fiber.Ctx, obj interface{}) error {
	res := extmodel.ApiResponse{
		StatusCode: 500,
		Success:    false,
		Result:     obj,
	}
	return c.Status(500).JSON(res)
}

func Unauthorized(c *fiber.Ctx) error {
	res := extmodel.ApiResponse{
		StatusCode: 401,
		Success:    false,
		Result:     "Unauthorized",
	}
	return c.Status(401).JSON(res)
}
