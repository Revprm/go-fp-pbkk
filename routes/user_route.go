package routes

import (
	"github.com/Revprm/go-fp-pbkk/constants"
	"github.com/Revprm/go-fp-pbkk/controller"
	"github.com/Revprm/go-fp-pbkk/middleware"
	"github.com/Revprm/go-fp-pbkk/service"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	routes := route.Group("/api/user")
	{
		// User
		routes.POST("", userController.Register)
		routes.POST("/login", userController.Login)
		routes.GET("", middleware.Authenticate(jwtService), middleware.OnlyAllow(constants.ENUM_ROLE_ADMIN), userController.GetAllUser)
		routes.DELETE("", middleware.Authenticate(jwtService), userController.Delete)
		routes.PATCH("", middleware.Authenticate(jwtService), userController.Update)
		routes.GET("/me", middleware.Authenticate(jwtService), userController.Me)
		routes.POST("/verify_email", userController.VerifyEmail)
		routes.POST("/send_verification_email", userController.SendVerificationEmail)
	}
}
