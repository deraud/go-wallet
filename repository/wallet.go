package repository

import (
	"main/models"

	"gorm.io/gorm"
)

type WalletRepository interface {
	FindByUserID(userID string) (*models.Wallet, error)
	Create(wallet *models.Wallet) error
	Update(wallet *models.Wallet) error
	CreateTransaction(transaction *models.Transaction) error
}

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{db: db}
}

func (r *walletRepository) FindByUserID(userID string) (*models.Wallet, error) {
	var wallet models.Wallet
	err := r.db.Where("user_id = ?", userID).First(&wallet).Error
	return &wallet, err
}

func (r *walletRepository) Create(wallet *models.Wallet) error {
	return r.db.Create(wallet).Error
}

func (r *walletRepository) Update(wallet *models.Wallet) error {
	return r.db.Save(wallet).Error
}

func (r *walletRepository) CreateTransaction(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}
