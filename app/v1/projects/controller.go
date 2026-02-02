package projects

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Success",
		"data": []map[string]interface{}{
			{
				"project_name": "High-Performance Search Engine",
				"problem": "Legacy search system couldn't handle 1M+ daily queries with acceptable latency.",
				"solution": "Built a distributed search service using Golang and Apache Solr with custom ranking algorithms and intelligent caching.",
				"Tech Stack": []string{"Golang", "Apache Solr", "Redis", "Kubernetes", "gRPC"},
				"Impact": []string{
					"Reduced P99 latency from 800ms to 120ms",
					"Handles 5000+ requests/second",
					"99.99% uptime achieved",
				},
			},
			{
				"project_name": "API Gateway & Rate Limiter",
				"problem": "Microservices lacked unified authentication and were vulnerable to traffic spikes.",
				"solution": "Designed and implemented a custom API gateway with JWT auth, rate limiting, and circuit breaker patterns.",
				"Tech Stack": []string{"Golang", "PostgreSQL", "Redis", "Docker", "Prometheus"},
				"Impact": []string{
					"Unified auth across 15+ services",
					"Blocked 100K+ malicious requests daily",
					"Reduced API abuse by 95%",
				},
			},
			{
				"project_name": "Real-Time Data Pipeline",
				"problem": "Batch processing created 4-hour delays in business analytics dashboards.",
				"solution": "Architected an event-driven pipeline using message queues for real-time data ingestion and processing.",
				"Tech Stack": []string{"Golang", "RabbitMQ", "PostgreSQL", "Grafana", "AWS"},
				"Impact": []string{
					"Near real-time data availability",
					"Processes 2M+ events daily",
					"Reduced operational costs by 30%",
				},
			},
			{
				"project_name": "Auto-Scaling Infrastructure",
				"problem": "Manual scaling during traffic spikes caused service degradation and high costs.",
				"solution": "Implemented Kubernetes-based auto-scaling with custom metrics and predictive scaling algorithms.",
				"Tech Stack": []string{"Kubernetes", "Terraform", "AWS EKS", "Prometheus", "Golang"},
				"Impact": []string{
					"Zero downtime during 10x traffic spikes",
					"40% reduction in cloud costs",
					"Self-healing infrastructure",
				},
			},
		},
	})
}