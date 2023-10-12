package app

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// Define an endpoint to fetch all data
	r.GET("/api/data", nil)

	return r
}
