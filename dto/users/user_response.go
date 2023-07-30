package usersdto

import "time"

type UserResponse struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"fullname" form:"fullname"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	Phone     string    `json:"phone" form:"phone"`
	Address   string    `json:"address" form:"address"`
	Subscribe string    `json:"subscribe" form:"subscribe"`
	Role      string    `json:"role"`
	DueDate   time.Time `json:"due_date" form:"due_date"`
}
