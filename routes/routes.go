package routes

import (
    "github.com/gin-gonic/gin"
    "backend-go/controllers"
    "backend-go/middleware"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/login", controllers.Login)

    auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())
    {
        auth.GET("/schedules", controllers.GetSchedules)
        auth.POST("/schedules", controllers.CreateSchedule)
        auth.PUT("/schedules/:id", controllers.UpdateSchedule)
        auth.DELETE("/schedules/:id", controllers.DeleteSchedule)
    }
}
