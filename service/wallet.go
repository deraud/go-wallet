package service

import (
	"errors"
	"main/models"
	"main/repository"

	"gorm.io/gorm"
)

type WalletService interface {
	GetBalance(userID string) (*models.Wallet, error)
	Withdraw(userID string, amount float64) error
}

type walletService struct {
	repo repository.WalletRepository
	db   *gorm.DB
}

func NewWalletService(repo repository.WalletRepository, db *gorm.DB) WalletService {
	return &walletService{
		repo: repo,
		db:   db,
	}
}

func (s *walletService) GetBalance(userID string) (*models.Wallet, error) {
	wallet, err := s.repo.FindByUserID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("wallet not found")
		}
		return nil, err
	}
	return wallet, nil
}

func (s *walletService) Withdraw(userID string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	wallet, err := s.repo.FindByUserID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			tx.Rollback()
			return errors.New("wallet not found")
		}
		tx.Rollback()
		return err
	}

	if wallet.Balance < amount {
		tx.Rollback()
		return errors.New("insufficient balance")
	}

	balanceBefore := wallet.Balance
	wallet.Balance -= amount

	if err := s.repo.Update(wallet); err != nil {
		tx.Rollback()
		return err
	}

	transaction := &models.Transaction{
		WalletID:        wallet.ID,
		TransactionType: "withdraw",
		Amount:          amount,
		BalanceBefore:   balanceBefore,
		BalanceAfter:    wallet.Balance,
		Description:     "Withdrawal",
	}

	if err := s.repo.CreateTransaction(transaction); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}