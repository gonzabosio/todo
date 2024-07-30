package main

import (
	"net/http"
	"todo/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	routes.ConnectDB()

	e := echo.New()

	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "Tasks API") })
	e.GET("/task", routes.GetTasks)
	e.POST("/task", routes.PostTask)
	e.PATCH("/task/:id", routes.PatchTask)
	e.DELETE("/task/:id", routes.DeleteTask)

	e.Logger.Fatal(e.Start(":3000"))
}
