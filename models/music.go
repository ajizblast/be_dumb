package models

type Music struct {
	ID       int    `json:"id" form:"id" gorm:"auto_increment:primary_key"`
	Title    string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Year     int    `json:"year" form:"year"`
	Image    string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Song     string `json:"song" form:"song" gorm:"type: varchar(255)"`
	ArtistID int    `json:"artist_id" form:"artist_id" gorm:"type: int"`
	Artist   Artist `json:"artist"`
}

type MusicResponse struct {
	ID        int    `json:"id" form:"id"`
	Title     string `form:"title"`
	Year      int    `form:"year"`
	Thumbnail string `form:"thumbnail"`
	Song      string `form:"song"`
	ArtistID  int    `form:"artist_id"`
	Artist    Artist `json:"artist" `
}

func (MusicResponse) TableName() string {
	return "Music"
}
