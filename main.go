package main

import (
	"portfolio-backend/config"
	"portfolio-backend/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	config.SetupRouter(r)

	r.Run(":8080")
}