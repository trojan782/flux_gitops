package main

import (
    "github.com/gin-gonic/gin"
    "go-api/internal/handlers"
    "go-api/internal/models"
    "log"
)

func main() {
    // Initialize store
    store := models.NewTaskStore()
    
    // Initialize handlers
    handler := handlers.NewHandler(store)
    
    // Initialize router
    r := gin.Default()
    
    // Routes
    r.GET("/health", handler.HealthCheck)
	r.GET("/", handler.GetHostInfo)
    
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