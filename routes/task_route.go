package routes

import (
	"github.com/Revprm/go-fp-pbkk/constants"
	"github.com/Revprm/go-fp-pbkk/controller"
	"github.com/Revprm/go-fp-pbkk/middleware"
	"github.com/Revprm/go-fp-pbkk/service"
	"github.com/gin-gonic/gin"
)

func Task(route *gin.Engine, taskController controller.TaskController, jwtService service.JWTService) {
	routes := route.Group("/api/task")
	{
		// Tasks
		routes.POST("", middleware.Authenticate(jwtService), taskController.CreateTask)
		routes.GET("", middleware.Authenticate(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_ADMIN), taskController.GetTasksWithPagination)
		routes.GET("/:id", middleware.Authenticate(jwtService), taskController.GetTaskByID)
		routes.PATCH("/:id", middleware.Authenticate(jwtService), taskController.UpdateTask)
		routes.DELETE("/:id", middleware.Authenticate(jwtService), taskController.DeleteTask)
	}
}
