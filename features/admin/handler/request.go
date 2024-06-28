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

type RegionCodeRequest struct {
	Code        string `json:"code"`
	Region      string `json:"region"`
	FullAddress string `json:"full_address"`
	PhoneNumber int    `json:"phone"`
	Price       int    `json:"price"`
	AdminID     uint   `json:"admin_id_perwakilan"`
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

func RequestToRegionCode(input RegionCodeRequest) admin.RegionCode {
	return admin.RegionCode{
		ID:          input.Code,
		Region:      input.Region,
		FullAddress: input.FullAddress,
		PhoneNumber: input.PhoneNumber,
		Price:       input.Price,
		AdminID:     input.AdminID,
	}
}

func generateID() uint {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Int63n(9999999999-1000000000) + 1000000000
	return uint(randomNumber)
}
