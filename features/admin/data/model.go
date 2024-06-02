package data

import (
	"jastip-jakarta/features/admin"

	"gorm.io/gorm"
)

type Admin struct {
	ID uint `gorm:"primaryKey" json:"id"`
	gorm.Model
	Name         string
	Email        string
	Password     string
	PhoneNumber  int
	PhotoProfile string
	Role         string
}

func AdminToModel(input admin.Admin) Admin {
	return Admin{
		ID:           input.ID,
		Name:         input.Name,
		Email:        input.Email,
		Password:     input.Password,
		PhoneNumber:  input.PhoneNumber,
		PhotoProfile: input.PhotoProfile,
		Role:         input.Role,
	}
}

func (u Admin) ModelToAdmin() admin.Admin {
	return admin.Admin{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		Password:     u.Password,
		PhoneNumber:  u.PhoneNumber,
		PhotoProfile: u.PhotoProfile,
		Role:         u.Role,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}
