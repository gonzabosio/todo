package main

import (
	"net/http"
	"todo/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	routes.ConnectDB()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
	}))

	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "Tasks API") })
	e.GET("/task", routes.GetTasks)
	e.POST("/task", routes.PostTask)
	e.PATCH("/task/:id", routes.PatchTask)
	e.DELETE("/task/:id", routes.DeleteTask)

	e.Logger.Fatal(e.Start(":3000"))
}
