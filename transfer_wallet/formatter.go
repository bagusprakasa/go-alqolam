package transferwallet

import "time"

type TransferWalletFormatter struct {
	ID           int       `json:"id"`
	FromWalletId int       `json:"from_wallet_id"`
	ToWalletId   int       `json:"to_wallet_id"`
	Total        float64   `json:"total"`
	Date         time.Time `json:"date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FormatTransferWallet(transferWallet TransferWallet) TransferWalletFormatter {
	formatter := TransferWalletFormatter{
		ID:           transferWallet.ID,
		FromWalletId: transferWallet.FromWalletId,
		ToWalletId:   transferWallet.ToWalletId,
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
