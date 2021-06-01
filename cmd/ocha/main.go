/*
 * Ocha - An TinyURL Alternative
 */
package main

import (
	_ "github.com/rahulvramesh/ocha-url/api"
	"github.com/rahulvramesh/ocha-url/internal/app/ocha"
	"github.com/rahulvramesh/ocha-url/internal/pkg/bigcache"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Ocha URL Shorten Service
// @version 1.0
// @description Ocha URL Shorten Service

// @contact.name rahulvramesh
// @contact.email rahulvramesh@hotmail.com

// @host localhost:8080
// @BasePath /
func main() {

	// Init Routes
	r := ocha.InitRoutes()
	// Initialize Store
	bigcache.InitializeDataSource()

	//Swagger Setup
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()

}
