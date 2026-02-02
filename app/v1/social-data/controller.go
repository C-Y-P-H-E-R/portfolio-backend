package social_data

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Success",
		"data": map[string]interface{}{
			"linkedin": map[string]string{
				"url": "https://www.linkedin.com/in/kushagrasharmaofficial",
				"username": "Kushagra Sharma",
				"registered_email": "myselflks@yahoo.com",
			},
			"github": map[string]string{
				"url": "https://github.com/C-Y-P-H-E-R",
				"username": "C-Y-P-H-E-R",
				"registered_email": "quasarkid123@gmail.com",
			},
			"leetcode": map[string]string{
				"url": "https://leetcode.com/u/sololevelerIndia/",
				"username": "sololevelerIndia",
				"registered_email": "quasarkid123@gmail.com",
			},
			"mail": map[string]string{
				"url": "mailto:sde.kushagra@gmail.com",
				"username": "sde.kushagra",
				"registered_email": "sde.kushagra@gmail.com",
			},
		},
	})
}