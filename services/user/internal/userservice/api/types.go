package api

import "rideshare/services/user/internal/userservice/internal/db"

type UserRegistrationParam struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PublicUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewPublicUserFromInternal(u db.User) PublicUser {
	return PublicUser{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
