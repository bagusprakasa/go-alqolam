package incomecategories

import "time"

type IncomeCategory struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Target    float64   `json:"target"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
