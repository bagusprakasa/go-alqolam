package user

import "time"

type UserFormatter struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	Role            string    `json:"role"`
	RememberToken   string    `json:"remember_token"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		EmailVerifiedAt: user.EmailVerifiedAt,
		RememberToken:   token,
		Role:            user.Role,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		DeletedAt:       user.DeletedAt,
	}

	return formatter
}

func FormatUserNonToken(user User) UserFormatter {
	formatter := UserFormatter{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		EmailVerifiedAt: user.EmailVerifiedAt,
		Role:            user.Role,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		DeletedAt:       user.DeletedAt,
	}

	return formatter
}

func FormatUsers(user []User) []UserFormatter {
	usersFormatter := []UserFormatter{}

	for _, user := range user {
		userFormatter := FormatUserNonToken(user)
		usersFormatter = append(usersFormatter, userFormatter)
	}

	return usersFormatter
}
