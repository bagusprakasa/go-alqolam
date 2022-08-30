package wallet

type WalletInput struct {
	Type          string `json:"type" binding:"required"`
	Name          string `json:"name" binding:"required"`
	AccountNumber string `json:"account_number"`
}

type GetDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
