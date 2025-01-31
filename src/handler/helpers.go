package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UnifiedResponse defines the standard response structure
type UnifiedResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // Optional field for data
}

// Response sends a standardized JSON response
func Response(c *gin.Context, statusCode int, status string, message string, data interface{}) {
	c.JSON(statusCode, UnifiedResponse{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func ResponseWithMsg(c *gin.Context, statusCode int, status string, message string) {
	c.JSON(statusCode, UnifiedResponse{
		Status:  status,
		Message: message,
	})
}

func ResponseWithError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, UnifiedResponse{
		Status:  "error",
		Message: message,
		Data:    nil,
	})
}

func ResponseNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, UnifiedResponse{
		Status:  "error",
		Message: message,
		Data:    nil,
	})
}
