package controllers

import (
	"CRUD/src/domain"
	"CRUD/src/interfaces/database"
	"CRUD/src/usecase"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c echo.Context) {
	u := domain.User{}
	c.Bind(&u)
	controller.Interactor.Add(u)
	createdUsers := controller.Interactor.GetInfo()
	c.JSON(201, createdUsers)
	return
}
func (controller *UserController) GetUsers() []domain.User {
	res := controller.Interactor.GetInfo()
	return res
}
func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}
