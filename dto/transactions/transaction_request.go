package transactiondto

import (
	"BE-Dumbsound/models"
	"time"
)

type CreateTransactionRequest struct {
	UserID        int         `json:"userid"`
	User          models.User `json:"user"`
	StartDate     time.Time   `json:"start_date"`
	DueDate       time.Time   `json:"due_date"`
	Price         int64       `json:"price"`
	StatusUser    string      `json:"status_user" gorm:"type: varchar(255)"`
	StatusPayment string      `json:"status_payment" gorm:"type: varchar(255)"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}
