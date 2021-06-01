// Package handlers - common handles for services
package handlers

import "github.com/gin-gonic/gin"

// PingHandler - ping handler
func PingHandler(c *gin.Context){
	c.String(200, "pong")
}


