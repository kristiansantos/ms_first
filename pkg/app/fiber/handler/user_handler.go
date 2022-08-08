package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kristiansantos/ms_first/internal/dto"
	"github.com/kristiansantos/ms_first/internal/repository"
	"github.com/kristiansantos/ms_first/internal/service"
)

type UserHandler interface {
	LogIn(ctx *fiber.Ctx) error
	LogOut(ctx *fiber.Ctx) error
	Index(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type userHandler struct {
	indexService  service.IndexUserServicer
	showService   service.ShowUserServicer
	createService service.CreateUserServicer
	updateService service.UpdateUserServicer
	deleteService service.DeleteUserServicer
}

func NewUserHandler(repository repository.UserRepository) UserHandler {
	return &userHandler{
		indexService:  service.NewUserIndexService(repository),
		showService:   service.NewUserShowService(repository),
		createService: service.NewUserCreateService(repository),
		updateService: service.NewUserUpdateService(repository),
		deleteService: service.NewUserDeleteService(repository),
	}
}

func (u *userHandler) LogIn(c *fiber.Ctx) error {
	fmt.Println(c)
	// Create database connection.
	return nil
}

func (u *userHandler) LogOut(c *fiber.Ctx) error {
	fmt.Println(c)
	// Create database connection.
	return nil
}

func (u *userHandler) Index(c *fiber.Ctx) error {
	users, err := u.indexService.Execute(nil)

	if err != nil {
		return c.
			Status(http.StatusInternalServerError).
			JSON(nil)
	}

	return c.
		Status(http.StatusOK).
		JSON(users)
}

func (u *userHandler) Show(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := u.showService.Execute(id)

	if err != nil {
		return c.
			Status(http.StatusNotFound).
			JSON(nil)
	}

	return c.
		Status(http.StatusOK).
		JSON(user)
}

func (u *userHandler) Create(c *fiber.Ctx) error {
	var userCreate dto.UserCreate

	if err := c.BodyParser(&userCreate); err != nil {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(nil)
	}

	user, err := u.createService.Execute(userCreate)

	if err != nil {
		return c.
			Status(http.StatusNotFound).
			JSON(nil)
	}

	return c.
		Status(http.StatusOK).
		JSON(user)
}

func (u *userHandler) Update(c *fiber.Ctx) error {
	var userUpdate dto.UserUpdate
	id := c.Params("id")

	if err := c.BodyParser(&userUpdate); err != nil {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(nil)
	}

	user, err := u.updateService.Execute(id, userUpdate)

	if err != nil {
		return c.
			Status(http.StatusNotFound).
			JSON(nil)
	}

	return c.
		Status(http.StatusOK).
		JSON(user)
}

func (u *userHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := u.deleteService.Execute(id)

	if err != nil {
		return c.
			Status(http.StatusNotFound).
			JSON(nil)
	}

	return c.
		Status(http.StatusNoContent).
		JSON(nil)
}
