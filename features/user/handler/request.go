package handler

import (
	"jastip-jakarta/features/user"
	"math/rand"
	"mime/multipart"
	"time"
)

type UserRequest struct {
	ID          uint
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber int    `json:"phone" form:"phone"`
}

type UserUpdateRequest struct {
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	PhoneNumber  int    `json:"phone" form:"phone"`
	PhotoProfile string `json:"photo_profile" form:"photo_profile"`
}

type LoginRequest struct {
	EmailOrPhone string `json:"email_or_phone" form:"email_or_phone"`
	Password     string `json:"password" form:"password"`
}

func RequestToUser(input UserRequest) user.User {
	return user.User{
		ID:          generateID(),
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
	}
}

func UpdateRequestToUser(input UserUpdateRequest, fileHeader *multipart.FileHeader) user.User {
	return user.User{
		Name:         input.Name,
		Email:        input.Email,
		Password:     input.Password,
		PhoneNumber:  input.PhoneNumber,
		PhotoProfile: fileHeader,
	}
}

func generateID() uint {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Int63n(9999999999-1000000000) + 1000000000
	return uint(randomNumber)
}
