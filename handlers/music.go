package handlers

import (
	musicdto "BE-Dumbsound/dto/music"
	dto "BE-Dumbsound/dto/result"
	"BE-Dumbsound/models"
	"BE-Dumbsound/repositories"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerMusic struct {
	MusicRepository repositories.MusicRepository
}

func HandleMusic(MusicRepository repositories.MusicRepository) *handlerMusic {
	return &handlerMusic{MusicRepository}
}
func (h *handlerMusic) FindMusics(c echo.Context) error {
	thumbnail, err := h.MusicRepository.FindMusics()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	response := dto.SuccesResult{Code: http.StatusOK, Data: thumbnail}
	return c.JSON(http.StatusOK, response)
}

func (h *handlerMusic) GetMusic(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	thumbnail, err := h.MusicRepository.GetMusic(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: thumbnail})
}

func (h *handlerMusic) CreateMusic(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	dataSong := c.Get("dataSong").(string)

	artistID, _ := strconv.Atoi(c.FormValue("artist_id"))
	year, _ := strconv.Atoi(c.FormValue("year"))
	request := musicdto.MusicRequest{
		Title:    c.FormValue("title"),
		Year:     year,
		Image:    dataFile,
		Song:     dataSong,
		ArtistID: artistID,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, _ := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "dumb sound"})
	respAudio, _ := cld.Upload.Upload(ctx, dataSong, uploader.UploadParams{Folder: "dumb sound"})

	if err != nil {
		fmt.Println(err.Error())
	}

	music := models.Music{
		Title:    request.Title,
		Year:     request.Year,
		Image:    resp.SecureURL,
		Song:     respAudio.SecureURL,
		ArtistID: request.ArtistID,
	}

	data, err := h.MusicRepository.CreateMusic(music)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	// data2, err := h.MusicRepository.GetMusic(data.ID)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	// }

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func (h *handlerMusic) DeleteMusic(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	trip, err := h.MusicRepository.GetMusic(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.MusicRepository.DeleteMusic(trip, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}
