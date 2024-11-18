package main

import (
	"go-api/internal/handlers"
	"go-api/internal/models"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// metrics
var (
    httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Total number of HTTP requests",
    }, []string{"method", "endpoint", "status"})

    httpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
        Name: "http_request_duration_seconds",
        Help: "Duration of HTTP requests",
    }, []string{"method", "endpoint"})
)


// middle ware

func prometheusMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        
        c.Next()
        
        status := strconv.Itoa(c.Writer.Status())
        duration := time.Since(start)
        
        httpRequestsTotal.WithLabelValues(c.Request.Method, c.FullPath(), status).Inc()
        httpRequestDuration.WithLabelValues(c.Request.Method, c.FullPath()).Observe(duration.Seconds())
    }
}
func main() {
    // Initialize store
    store := models.NewTaskStore()
    
    // Initialize handlers
    handler := handlers.NewHandler(store)
    
    // Initialize router
    r := gin.Default()

	r.Use(prometheusMiddleware())
    
    // Routes
    r.GET("/health", handler.HealthCheck)
	r.GET("/", handler.GetHostInfo)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
    
    // API v1 group
    v1 := r.Group("/api/v1")
    {
        tasks := v1.Group("/tasks")
        {
            tasks.GET("", handler.GetTasks)
            tasks.GET("/:id", handler.GetTask)
            tasks.POST("", handler.CreateTask)
            tasks.PUT("/:id", handler.UpdateTask)
            tasks.DELETE("/:id", handler.DeleteTask)
        }
    }
    
    // Start server
    log.Fatal(r.Run(":8080"))
}