package member

import (
	"go-alqolam/region"
	"time"
)

type Member struct {
	ID        int       `json:"id" gorm:"primary_key"`
	RegionID  int       `json:"region_id"`
	Name      string    `json:"name"`
	Phone     string    `json:"Phone"`
	Address   string    `json:"Address"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Region    region.Region
}
