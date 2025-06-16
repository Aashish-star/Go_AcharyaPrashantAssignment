package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse is a generic structure for API error responses
type ErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorHandler is a middleware that recovers from panics and handles errors
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if any errors occurred during request handling
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
		}
	}
}
