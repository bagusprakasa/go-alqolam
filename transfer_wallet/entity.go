package transferwallet

import (
	"go-alqolam/wallet"
	"time"
)

type TransferWallet struct {
	ID           int       `json:"id" gorm:"primary_key"`
	FromWalletId int       `json:"from_wallet_id"`
	ToWalletId   int       `json:"to_wallet_id"`
	Total        float64   `json:"total"`
	Date         time.Time `json:"date"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	FromWallet   wallet.Wallet
	ToWallet     wallet.Wallet
}
