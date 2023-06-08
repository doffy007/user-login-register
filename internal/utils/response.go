package utils

import "github.com/gin-gonic/gin"

func ResponseSuccessGet(data interface{}) gin.H {
	return gin.H{
		"message": "Success",
		"data":    data,
	}
}

func ResponseBaseSuccess(message string) gin.H {
	return gin.H{
		"message": message,
	}
}

func ResponseBaseError2(message string) gin.H {
	return gin.H{
		"message": message,
		"data":    nil,
	}
}

func NewErrorResponse(c *gin.Context, statusCode int, message string, err interface{}) {
	r := struct {
		StatusCode int         `json:"status_code"`
		Message    string      `json:"message"`
		Error      interface{} `json:"errors"`
	}{
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	}

	c.JSON(statusCode, r)
}

func NewSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	r := struct {
		StatusCode int         `json:"status_code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
	}{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	c.JSON(statusCode, r)
}
