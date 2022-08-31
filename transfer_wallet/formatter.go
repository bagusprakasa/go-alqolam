package transferwallet

import "time"

type TransferWalletFormatter struct {
	ID           int       `json:"id"`
	FromWalletId int       `json:"from_wallet_id"`
	ToWalletId   int       `json:"to_wallet_id"`
	FromWallet   string    `json:"from_wallet"`
	ToWallet     string    `json:"to_wallet"`
	Total        float64   `json:"total"`
	Date         time.Time `json:"date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FormatTransferWallet(transferWallet TransferWallet) TransferWalletFormatter {
	fromWallet := transferWallet.FromWallet
	toWallet := transferWallet.ToWallet
	formatter := TransferWalletFormatter{
		ID:           transferWallet.ID,
		FromWalletId: transferWallet.FromWalletId,
		FromWallet:   fromWallet.Name,
		ToWalletId:   transferWallet.ToWalletId,
		ToWallet:     toWallet.Name,
		Total:        transferWallet.Total,
		Date:         transferWallet.Date,
		CreatedAt:    transferWallet.CreatedAt,
		UpdatedAt:    transferWallet.UpdatedAt,
	}

	return formatter
}

func FormatTransferWallets(transferWallet []TransferWallet) []TransferWalletFormatter {
	transferWalletsFormatter := []TransferWalletFormatter{}

	for _, transferWallet := range transferWallet {
		transferWalletFormatter := FormatTransferWallet(transferWallet)
		transferWalletsFormatter = append(transferWalletsFormatter, transferWalletFormatter)
	}

	return transferWalletsFormatter
}
