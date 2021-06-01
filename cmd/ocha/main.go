/*
 * Ocha - An TinyURL Alternative
 */
package main

import (
	"github.com/rahulvramesh/ocha-url/internal/app/ocha"
	"github.com/rahulvramesh/ocha-url/internal/pkg/bigcache"
)

// @title Ocha URL Shorten Service
// @version 1.0
// @description Ocha URL Shorten Service

// @contact.name rahulvramesh
// @contact.email rahulvramesh@hotmail.com

// @host localhost:8000
// @BasePath /
func main() {

	// Init Routes
	r := ocha.InitRoutes()

	// Initialize Store
	bigcache.InitializeDataSource()

	r.Run()

}
