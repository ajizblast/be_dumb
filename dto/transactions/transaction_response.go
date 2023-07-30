package transactiondto

import (
	"BE-Dumbsound/models"
	"time"
)

type TransactionResponse struct {
	ID            int         `json:"id"`
	UserID        int         `json:"userid"`
	User          models.User `json:"user"`
	Price         int64       `json:"price"`
	StartDate     time.Time   `json:"start_date" gorm:"type: varchar(255)"`
	DueDate       time.Time   `json:"due_date" gorm:"type: varchar(255)"`
	StatusUser    string      `json:"status_user" gorm:"type: varchar(255)"`
	StatusPayment string      `json:"status_payment" gorm:"type: varchar(255)"`
}
