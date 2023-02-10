package infrastructure

import (
	controllers "CRUD/src/interfaces/api"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Init() {
	e := echo.New()
	cfg := GetConfig()
	fmt.Println("Data from init", cfg)
	userController := controllers.NewUserController(NewSqlHandler(cfg.Storage))
	e.GET("/users", func(c echo.Context) error {
		users := userController.GetUsers()
		c.Bind(&users)
		return c.JSON(http.StatusOK, users)
	})
	e.POST("/users", func(c echo.Context) error {
		userController.Create(c)
		return c.String(http.StatusOK, "created")
	})
	e.DELETE("users/:id", func(c echo.Context) error {
		id := c.Param("id")
		userController.Delete(id)
		return c.String(http.StatusOK, "deleted")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
