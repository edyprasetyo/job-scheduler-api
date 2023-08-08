package response_handler

import (
	"jobschedulerapi/domain/extmodel"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, obj interface{}) {
	res := extmodel.ApiResponse{
		StatusCode: 200,
		Success:    true,
		Result:     obj,
	}
	c.JSON(200, res)
}

func ValidationError(c *gin.Context, obj interface{}) {
	res := extmodel.ApiResponse{
		StatusCode: 400,
		Success:    false,
		Result:     obj,
	}
	c.JSON(400, res)
}

func InternalError(c *gin.Context, obj interface{}) {
	res := extmodel.ApiResponse{
		StatusCode: 500,
		Success:    false,
		Result:     obj,
	}
	c.JSON(500, res)
}

func Unauthorized(c *gin.Context) {
	res := extmodel.ApiResponse{
		StatusCode: 401,
		Success:    false,
		Result:     "Unauthorized",
	}
	c.JSON(401, res)
}
