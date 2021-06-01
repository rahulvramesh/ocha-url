package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rahulvramesh/ocha-url/internal/app/ocha/models"
	"github.com/rahulvramesh/ocha-url/internal/pkg/httputil"
	_ "github.com/rahulvramesh/ocha-url/internal/pkg/httputil"
	"github.com/rahulvramesh/ocha-url/internal/pkg/services"
)

// CreateShortURLHandler - create short url handler
// @Summary Create short url handler
// @Description Create short url handler
// @ID create-short-url
// @Accept  json
// @Produce  json
// @Param short-url body models.CreateRequest true "New Short URL"
// @Success 201 {object} models.CreateResponse
// @Failure 400,409,422 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.HTTPError
// @Router /shorten [post]
func CreateShortURLHandler(ctx *gin.Context) {
	var (
		req models.CreateRequest
	)

	err := ctx.BindJSON(&req)
	if err != nil {
		httputil.NewError(ctx,400, errors.New("bad request"))
	}

	// Check If Key Exists
	svc := services.OchaService{}
	err = svc.CheckIfKeyExists(req.Key)
	if err != nil {
		// Response with key exists error
	}

	ctx.JSON(200, req)
}

