package routes

import (
	"BE-Dumbsound/handlers"
	"BE-Dumbsound/pkg/middleware"
	"BE-Dumbsound/pkg/mysql"
	"BE-Dumbsound/repositories"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/check-auth", middleware.Auth(h.CheckAuth))
}
