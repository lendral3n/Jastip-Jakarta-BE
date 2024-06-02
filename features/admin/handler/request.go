package handler

import (
	"jastip-jakarta/features/admin"
	"math/rand"
	"time"
)

type AdminRequest struct {
	ID          uint
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber int    `json:"phone" form:"phone"`
	Role        string `json:"role" form:"role"`
}

type LoginRequest struct {
	EmailOrPhone string `json:"email_or_phone" form:"email_or_phone"`
	Password     string `json:"password" form:"password"`
}

func RequestToAdmin(input AdminRequest) admin.Admin {
	return admin.Admin{
		ID:          generateID(),
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
		Role:        input.Role,
	}
}

func generateID() uint {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Int63n(9999999999-1000000000) + 1000000000
	return uint(randomNumber)
}
