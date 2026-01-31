package config

import (
	"portfolio-backend/app/v1/health"
	"portfolio-backend/app/v1/social-data"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	v1 := r.Group("/go/api/v1")
	{
		v1.GET("/ping", health.Ping)
		v1.GET("/socials", social_data.GET)
	}
}