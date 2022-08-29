package user

import "time"

type User struct {
	ID              int       `json:"id" gorm:"primary_key"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	EmailVerifiedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"email_verified_at"`
	Password        string    `json:"password"`
	Role            string    `json:"role"`
	RememberToken   string    `json:"remember_token"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
