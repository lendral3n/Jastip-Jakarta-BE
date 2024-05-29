package data

import (
	"jastip-jakarta/features/user"

	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`
	gorm.Model
	Name         string
	Email        string
	Password     string
	PhoneNumber  int
	PhotoProfile string
}

func UserToModel(input user.User) User {
	photoProfileURL, ok := input.PhotoProfile.(string)
	if !ok {
		photoProfileURL = ""
	}

	return User{
		ID:           input.ID,
		Name:         input.Name,
		Email:        input.Email,
		Password:     input.Password,
		PhoneNumber:  input.PhoneNumber,
		PhotoProfile: photoProfileURL,
	}
}

func (u User) ModelToUser() user.User {
	return user.User{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		PhoneNumber:  u.PhoneNumber,
		PhotoProfile: u.PhotoProfile,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}
