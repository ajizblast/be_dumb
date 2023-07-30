package routes

import (
	"BE-Dumbsound/handlers"
	"BE-Dumbsound/pkg/middleware"
	"BE-Dumbsound/pkg/mysql"
	"BE-Dumbsound/repositories"

	"github.com/labstack/echo/v4"
)

func MusicRoutes(e *echo.Group) {
	musicRepository := repositories.RepositoryMusic(mysql.DB)
	h := handlers.HandleMusic(musicRepository)

	e.GET("/musics", h.FindMusics)
	e.GET("/music/:id", h.GetMusic)
	e.POST("/music", middleware.Auth(middleware.UploadFile(middleware.UploadSong(h.CreateMusic))))
}
