package models

import "time"

type User struct {
	ID         int       `json:"id" form:"id" gorm:"PRIMARY_KEY:AUTO_INCREMENT"`
	Fullname   string    `json:"fullname" form:"fullname" gorm:"type: varchar(255)"`
	Email      string    `json:"email" form:"email" gorm:"type: varchar(255)"`
	Password   string    `json:"password" form:"password" gorm:"type: varchar(255)"`
	Gender     string    `json:"gender" form:"gender" gorm:"type: varchar(255)"`
	Phone      string    `json:"phone" form:"phone" gorm:"type: varchar(255)"`
	Address    string    `json:"address" form:"address" gorm:"type: varchar(255)"`
	Subscribe  string    `json:"subscribe" form:"subscribe" gorm:"type: varchar(255)"`
	StatusUser string    `json:"status_user" form:"status_user" gorm:"type: varchar(255)"`
	Role       string    `json:"role" gorm:"type: varchar(255)"`
	DueDate    time.Time `json:"due_date" form:"due_date"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}
