package httputil

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// NewError - Build custom errors
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError -
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// KeyAlreadyExistsError - error if key already exits
func KeyAlreadyExistsError (ctx *gin.Context){
	NewError(ctx,409,errors.New("the the desired shortcode is already in use [shortcodes are case-sensitive]"))
}

// UnexpectedError - 500 error
func UnexpectedError (ctx *gin.Context){
	NewError(ctx,500,errors.New("unexpected error occurred"))
}

// KeyNotFoundError - key not available on the system
func KeyNotFoundError (ctx *gin.Context){
	NewError(ctx,404,errors.New("the shortcode cannot be found in the system"))
}