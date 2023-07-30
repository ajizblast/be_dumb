package musicdto

type MusicRequest struct {
	Title    string `json:"title" form:"title" validate:"required"`
	Year     int    `json:"year" form:"year" validate:"required"`
	Image    string `json:"image" form:"image"  validate:"required"`
	Song     string `form:"song" validate:"required"`
	ArtistID int    `json:"artist_id" form:"artist_id"  validate:"required"`
}
