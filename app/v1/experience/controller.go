package experience

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Success",
		"data": map[string]interface{}{
			"IndiaMart": map[string]interface{}{
				"position": "Software Engineer - Backend",
				"start_year": "2022",
				"end_year": "Present",
				"Location": "Noida, India",
				"description": "Building and scaling backend systems for India's largest B2B marketplace serving millions of businesses",
				"projects": []map[string]string{
					{
						"project_name": "Apache Solr Optimization",
						"project_description": "Optimized search indexing pipeline handling 10M+ documents, reduced index refresh time by 70% through sharding and caching strategies.",
					},
					{
						"project_name": "CI/CD Pipeline Enhancement",
						"project_description": "Implemented automated testing and deployment pipelines, reducing deployment time from hours to minutes with zero-downtime releases.",
					},
					{
						"project_name": "API Migration (PHP → Golang)",
						"project_description": "Led migration of critical search APIs from legacy PHP to Golang, achieving 60% reduction in response times and 40% cost savings on infrastructure.",
					},
					{
						"project_name": "Microservices Architecture",
						"project_description": "Designed and deployed 5+ microservices on Kubernetes, enabling independent scaling and 99.9% uptime across services.",
					},
				},
			},
		},
	})
}