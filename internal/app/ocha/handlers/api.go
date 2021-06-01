package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rahulvramesh/ocha-url/internal/app/ocha/models"
	"github.com/rahulvramesh/ocha-url/internal/pkg/httputil"
	_ "github.com/rahulvramesh/ocha-url/internal/pkg/httputil"
	"github.com/rahulvramesh/ocha-url/internal/pkg/services"
	"net/url"
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
		httputil.NewError(ctx, 400, errors.New("bad request"))
	}

	// Check If Key Exists
	svc := services.OchaService{}
	err = svc.CheckIfKeyExists(req.Key)
	if err == nil {
		// Response with key exists error
		httputil.KeyAlreadyExistsError(ctx)
		return
	}

	// Else Create The Entry
	err = svc.StoreLink(req.Key, req.URL)
	if err != nil {
		httputil.UnexpectedError(ctx)
		return
	}

	ctx.JSON(201, models.CreateResponse{Key: req.Key})
}

// GetByShortCodeHandler - Get URL Redirect By Shortcode
// @Summary Get URL Redirect By Shortcode
// @Description Get URL Redirect By Shortcode
// @ID get-short-url
// @Param shortcode path string true "Shortcode"
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.HTTPError
// @Router /{shortcode} [get]
func GetByShortCodeHandler(ctx *gin.Context) {

	// Decode Short Code
	shortcode,err := url.QueryUnescape(ctx.Param("shortcode"))
	if err != nil  {
		httputil.UnexpectedError(ctx)
		return
	}
	svc := services.OchaService{}

	err = svc.GetByKey(shortcode)
	if err != nil  {
		if svc.Data.URL != "" {
			httputil.UnexpectedError(ctx)
			return
		}
		httputil.KeyNotFoundError(ctx)
		return
	}
	ctx.Redirect(302,svc.Data.URL)
}


// GetStatusHandler - Get Status By Shortcode Handler
// @Summary Get Status By Shortcode Handler
// @Description Get Status By Shortcode Handler
// @ID get-short-url-status
// @Param shortcode path string true "Shortcode"
// @Success 200 {object} models.CreateResponse
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.HTTPError
// @Router /{shortcode}/stats [get]
func GetStatusHandler(ctx *gin.Context) {
	shortcode,err := url.QueryUnescape(ctx.Param("shortcode"))
	if err != nil  {
		httputil.UnexpectedError(ctx)
		return
	}
	svc := services.OchaService{}

	err = svc.GetByKey(shortcode)
	if err != nil  {
		if svc.Data.URL != "" {
			httputil.UnexpectedError(ctx)
			return
		}
		httputil.KeyNotFoundError(ctx)
		return
	}
	ctx.JSON(200,svc.Data)
}