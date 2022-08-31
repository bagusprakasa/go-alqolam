package transferwallet

import "time"

type TransferWalletInput struct {
	FromWalletId int       `json:"from_wallet_id" binding:"required"`
	ToWalletId   int       `json:"to_wallet_id" binding:"required"`
	Total        float64   `json:"total" binding:"required"`
	Date         time.Time `json:"date"`
}

type GetDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
