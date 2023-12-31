package routes

import (
	"BE-Dumbsound/handlers"
	"BE-Dumbsound/pkg/middleware"
	"BE-Dumbsound/pkg/mysql"
	"BE-Dumbsound/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", h.FindUsers)
	e.GET("/user/:id", h.GetUser)
	e.PATCH("/user/:id", middleware.Auth(h.UpdateUser))
}
