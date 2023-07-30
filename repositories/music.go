package repositories

import (
	"BE-Dumbsound/models"

	"gorm.io/gorm"
)

type MusicRepository interface {
	FindMusics() ([]models.Music, error)
	GetMusic(ID int) (models.Music, error)
	CreateMusic(trip models.Music) (models.Music, error)
	DeleteMusic(music models.Music, ID int) (models.Music, error)
}

func RepositoryMusic(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindMusics() ([]models.Music, error) {
	var Music []models.Music
	err := r.db.Preload("Artist").Find(&Music).Error

	return Music, err
}

func (r *repository) GetMusic(ID int) (models.Music, error) {
	var Music models.Music
	err := r.db.Preload("Artist").First(&Music, ID).Error

	return Music, err
}

func (r *repository) CreateMusic(music models.Music) (models.Music, error) {
	err := r.db.Preload("Artist").Create(&music).Error

	return music, err
}

func (r *repository) DeleteMusic(music models.Music, ID int) (models.Music, error) {
	err := r.db.Delete(&music, ID).Scan(&music).Error

	return music, err
}
