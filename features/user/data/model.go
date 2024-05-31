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
	return User{
		ID:           input.ID,
		Name:         input.Name,
		Email:        input.Email,
		Password:     input.Password,
		PhoneNumber:  input.PhoneNumber,
		PhotoProfile: input.PhotoProfile,
	}
}

func (u User) ModelToUser() user.User {
	return user.User{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		Password:     u.Password,
		PhoneNumber:  u.PhoneNumber,
		PhotoProfile: u.PhotoProfile,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}
