package handler

import (
	"jastip-jakarta/features/user"
	"jastip-jakarta/utils/time"
)

type UserResponse struct {
	ID           uint   `json:"user_id" form:"user_id"`
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	PhoneNumber  int    `json:"phone_number" form:"phone_number"`
	PhotoProfile string `json:"photo_profile" form:"photo_profile"`
	CreatedAt    string `json:"create_account"`
	UpdatedAt    string `json:"last_update"`
}

type UserResponseOrder struct {
	Name string `json:"name" form:"name"`
}

func UserToResponse(data *user.User) UserResponse {
	return UserResponse{
		ID:           data.ID,
		Name:         data.Name,
		Email:        data.Email,
		PhoneNumber:  data.PhoneNumber,
		PhotoProfile: data.PhotoProfile,
		CreatedAt:    time.FormatDateToIndonesian(data.CreatedAt),
		UpdatedAt:    time.FormatDateToIndonesian(data.UpdatedAt),
	}
}