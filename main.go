package main

import (
	"log"
	"os"

	"github.com/Revprm/go-fp-pbkk/command"
	"github.com/Revprm/go-fp-pbkk/config"
	"github.com/Revprm/go-fp-pbkk/controller"
	"github.com/Revprm/go-fp-pbkk/middleware"
	"github.com/Revprm/go-fp-pbkk/repository"
	"github.com/Revprm/go-fp-pbkk/routes"
	"github.com/Revprm/go-fp-pbkk/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	if len(os.Args) > 1 {
		flag := command.Commands(db)
		if !flag {
			return
		}
	}

	var (
		jwtService service.JWTService = service.NewJWTService()

		// Implementation Dependency Injection
		// Repository
		userRepository repository.UserRepository = repository.NewUserRepository(db)

		// Service
		userService service.UserService = service.NewUserService(userRepository, jwtService)

		// Controller
		userController controller.UserController = controller.NewUserController(userService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	// routes
	routes.User(server, userController, jwtService)
	server.Static("/assets", "./assets")

	// Check migration
	// if err := migrations.Seeder(db); err != nil {
	// 	log.Fatalf("error migration seeder: %v", err)
	// }

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	// Just in case it detects trailing slash
	server.RedirectTrailingSlash = true
	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
