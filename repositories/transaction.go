package repositories

import (
	"BE-Dumbsound/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(Transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, orderId int) (models.Transaction, error)
	DeleteTransaction(Transaction models.Transaction, ID int) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var Transaction []models.Transaction
	err := r.db.Preload("User").Find(&Transaction).Error

	return Transaction, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var Transaction models.Transaction
	err := r.db.Preload("User").First(&Transaction, ID).Error

	return Transaction, err
}

func (r *repository) CreateTransaction(Transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Film").Create(&Transaction).Error

	return Transaction, err
}

func (r *repository) DeleteTransaction(Transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&Transaction, ID).Scan(&Transaction).Error

	return Transaction, err
}

func (r *repository) UpdateTransaction(status string, orderId int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("User").First(&transaction, orderId)

	if status != transaction.StatusPayment && status == "success" {
		var user models.User
		r.db.First(&user, transaction.UserID)
		user.StatusUser = "active"
		user.DueDate = transaction.DueDate
		r.db.Save(&user)
	}

	transaction.StatusPayment = status
	err := r.db.Save(&transaction).Error
	return transaction, err
}
