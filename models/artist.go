package models

type Artist struct {
	ID          int    `json:"id" form:"id" gorm:"auto_increment:primary_key"`
	Name        string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Age         string `json:"age" form:"age" gorm:"type: varchar(255)"`
	Type        string `json:"type" form:"type" gorm:"type: varchar(255)"`
	StartCareer string `json:"start_career" form:"start_career" gorm:"type: varchar(255)"`
}
