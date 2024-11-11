package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "go-api/internal/models"
    "net/http"
	"os"
)

type Handler struct {
    store *models.TaskStore
}

func NewHandler(store *models.TaskStore) *Handler {
    return &Handler{store: store}
}

func (h *Handler) GetTasks(c *gin.Context) {
    tasks := h.store.GetTasks()
    c.JSON(http.StatusOK, tasks)
}

func (h *Handler) GetTask(c *gin.Context) {
    id := c.Param("id")
    task, exists := h.store.GetTask(id)
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, task)
}

func (h *Handler) CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    task.ID = uuid.New().String()
    task = h.store.CreateTask(task)
    c.JSON(http.StatusCreated, task)
}

func (h *Handler) UpdateTask(c *gin.Context) {
    id := c.Param("id")
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    task.ID = id
    updatedTask, exists := h.store.UpdateTask(task)
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    
    c.JSON(http.StatusOK, updatedTask)
}

func (h *Handler) DeleteTask(c *gin.Context) {
    id := c.Param("id")
    if deleted := h.store.DeleteTask(id); !deleted {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    c.Status(http.StatusNoContent)
}

func (h *Handler) HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": "healthy",
    })
}

func (h *Handler) GetHostInfo(c *gin.Context) {
    hostname, err := os.Hostname()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status": http.StatusInternalServerError,
            "message": "Error getting hostname",
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "message": "This app is running on host " + hostname,
    })
}