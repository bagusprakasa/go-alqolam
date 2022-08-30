package wallet

import "time"

type Wallet struct {
	ID            int       `json:"id" gorm:"primary_key"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	AccountNumber string    `json:"account_number"`
	Balance       float64   `json:"balance" default:"0"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
