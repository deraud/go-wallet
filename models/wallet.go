package models

import (
	"time"
)

type Wallet struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    string    `gorm:"unique;not null" json:"user_id"`
	Balance   float64   `gorm:"type:decimal(10,2);default:0" json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Transaction struct {
	ID              uint      `gorm:"primarykey" json:"id"`
	WalletID        uint      `gorm:"not null" json:"wallet_id"`
	TransactionType string    `gorm:"type:varchar(20);not null" json:"transaction_type"`
	Amount          float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	BalanceBefore   float64   `gorm:"type:decimal(10,2)" json:"balance_before"`
	BalanceAfter    float64   `gorm:"type:decimal(10,2)" json:"balance_after"`
	Description     string    `gorm:"type:text" json:"description"`
	CreatedAt       time.Time `json:"created_at"`
}
