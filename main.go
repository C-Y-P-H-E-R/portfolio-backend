package main

import (
	"portfolio-backend/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.SetupRouter(r)

	r.Run(":8080")
}