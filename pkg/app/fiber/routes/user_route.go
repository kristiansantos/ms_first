package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kristiansantos/ms_first/pkg/app/fiber/handler"
	mongoRepository "github.com/kristiansantos/ms_first/pkg/infra/mongodb"
	"github.com/kristiansantos/ms_first/pkg/mongodb"
)

// PrivateRoutes func for describe group of private routes.
func UserRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	ctx := context.TODO()

	db := mongodb.New(ctx)
	userRepository := mongoRepository.NewUsersRepository(db, ctx)
	userHandler := handler.NewUserHandler(userRepository)

	route.Get("/users", userHandler.Index)
	route.Get("/users/:id", userHandler.Show)

	route.Post("/users", userHandler.Create)

	route.Put("/users/:id", userHandler.Update)

	route.Delete("/users/:id", userHandler.Delete)

	route.Post("/signup", userHandler.LogIn)
	route.Post("/signin", userHandler.LogOut)
}
