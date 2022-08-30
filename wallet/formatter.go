package wallet

import "time"

type WalletFormatter struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	Balance       float64   `json:"balance"`
	AccountNumber string    `json:"account_number"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func FormatWallet(wallet Wallet) WalletFormatter {
	formatter := WalletFormatter{
		ID:            wallet.ID,
		Name:          wallet.Name,
		Type:          wallet.Type,
		Balance:       wallet.Balance,
		AccountNumber: wallet.AccountNumber,
		CreatedAt:     wallet.CreatedAt,
		UpdatedAt:     wallet.UpdatedAt,
	}

	return formatter
}

func FormatWallets(wallet []Wallet) []WalletFormatter {
	walletsFormatter := []WalletFormatter{}

	for _, wallet := range wallet {
		walletFormatter := FormatWallet(wallet)
		walletsFormatter = append(walletsFormatter, walletFormatter)
	}

	return walletsFormatter
}
