package ocha

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulvramesh/ocha-url/internal/app/ocha/handlers"
)

/*
 * Routes Collection
 */

func InitRoutes() *gin.Engine{
	r := gin.Default()

	r.GET("/ping",handlers.PingHandler)
	r.POST("/shorten",handlers.CreateShortURLHandler)
	r.GET("/:shortcode",handlers.GetByShortCodeHandler)
	r.GET("/:shortcode/stats",handlers.GetStatusHandler)

	return r
}