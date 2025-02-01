package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func ParseErrorMessage(err error) string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]string, len(ve))
		for i, fe := range ve {
			out[i] = GetErrorMessage(fe)
		}
		return strings.Join(out, " , ")
	}
	return err.Error()
}

func GetErrorMessage(fe validator.FieldError) string {

	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s field is required!", fe.Field())
	case "email":
		return fmt.Sprintf("Valid %s is required!", fe.Field())
	case "alpha":
		return fmt.Sprintf("%s field only character required!", fe.Field())
	case "oneof":
		return fmt.Sprintf("%s should be between %s!", fe.Field(), fe.Param())
	case "lte":
		return fmt.Sprintf("%s Should be less than %s", fe.Field(), fe.Param())
	case "gte":
		return fmt.Sprintf("%s Should be greater than %s", fe.Field(), fe.Param())
	}
	return "Unknown error"
}
