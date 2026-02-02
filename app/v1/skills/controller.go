package skills

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Success",
		"data": map[string]interface{}{
			"Backend": map[string]bool{
				"Golang": true,
				"PHP": true,
				"RESTful APIs": true,
				"gRPC": true,
				"Microservices": true,
				"Node": true,
				"Python": true,
				"Kafka": true,
			},
			"Databases": map[string]bool{
				"PostgreSQL": true,
				"OracleDB": true,
				"Redis": true,
				"MongoDB": true,
				"Solr": true,
			},
			"DevOps": map[string]bool{
				"Docker": true,
				"Kubernetes": true,
				"CI/CD": true,
				"Gitlab": true,
				"API Gateway": true,
			},
			"Cloud": map[string]bool{
				"EC2": true,
				"S3": true,
				"ALB": true,
				"EKS": true,
				"CloudWatch": true,
				"RDS": true,
				"Lambda": true,
			},
			"Tools & Frameworks": map[string]bool{
				"Git": true,
				"Linux": true,
				"Nginx": true,
				"RabbitMQ": true,
				"Prometheus": true,
				"Grafana": true,
				"Kibana": true,
			},
		},
	})
}