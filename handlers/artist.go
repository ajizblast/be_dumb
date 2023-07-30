package handlers

import (
	artistdto "BE-Dumbsound/dto/artists"
	dto "BE-Dumbsound/dto/result"
	"BE-Dumbsound/models"
	"BE-Dumbsound/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerArtist struct {
	ArtistRepository repositories.ArtistRepository
}

func HandleArtist(ArtistRepository repositories.ArtistRepository) *handlerArtist {
	return &handlerArtist{ArtistRepository}
}

func (h *handlerArtist) FindArtists(c echo.Context) error {
	artist, err := h.ArtistRepository.FindArtists()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: artist})
}

func (h *handlerArtist) GetArtist(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Invalid ID"})
	}

	artist, err := h.ArtistRepository.GetArtist(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: artist})
}

func (h *handlerArtist) CreateArtist(c echo.Context) error {
	get := new(artistdto.ArtistRequest)
	if err := c.Bind(get); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(get)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	artist := models.Artist{
		Name:        get.Name,
		Age:         get.Age,
		Type:        get.Type,
		StartCareer: get.StartCareer,
	}

	data, err := h.ArtistRepository.CreateArtist(artist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func (h *handlerArtist) DeleteArtist(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.ArtistRepository.GetArtist(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.ArtistRepository.DeleteArtist(user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}
