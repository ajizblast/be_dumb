package routes

import (
	"BE-Dumbsound/handlers"
	"BE-Dumbsound/pkg/middleware"
	"BE-Dumbsound/pkg/mysql"
	"BE-Dumbsound/repositories"

	"github.com/labstack/echo/v4"
)

func ArtistRoutes(e *echo.Group) {
	artistRepository := repositories.RepositoryArtist(mysql.DB)
	h := handlers.HandleArtist(artistRepository)

	e.GET("/artists", h.FindArtists)
	e.POST("/artists", middleware.Auth(h.CreateArtist))
	e.GET("/artists/:id", middleware.Auth(h.GetArtist))
	e.DELETE("/artists/:id", middleware.Auth(h.DeleteArtist))
}
