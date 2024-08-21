package handler

import (
	"fmt"
	"jastip-jakarta/features/admin"
	uh "jastip-jakarta/features/user"
	"math/rand"
	"sync"
	"time"
)

type AdminRequest struct {
	ID          uint
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber int    `json:"phone"`
	Role        string `json:"role"`
}

type LoginRequest struct {
	EmailOrPhone string `json:"email_or_phone"`
	Password     string `json:"password"`
}

type RegionCodeRequest struct {
	Code        string `json:"code"`
	Region      string `json:"region"`
	FullAddress string `json:"full_address"`
	PhoneNumber int    `json:"phone"`
	Price       int    `json:"price"`
	AdminID     uint   `json:"admin_id_perwakilan"`
}

type DeliveryBatchRequest struct {
	DeliveryBatch string
	Batch         int `json:"batch"`
	Year          int `json:"year"`
	Month         int `json:"month"`
}

type UserRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber int    `json:"phone" form:"phone"`
}

func RequestToUser(input UserRequest) uh.User {
	return uh.User{
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
	}
}

func RequestRegisterToUser(input UserRequest) uh.User {
	return uh.User{
		ID: generateIDUSer(),
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
	}
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

func RequestToDeliveryBatch(input DeliveryBatchRequest) admin.DeliveryBatch {
	id := fmt.Sprintf("%02d%04dB%d", input.Month, input.Year, input.Batch)
	return admin.DeliveryBatch{
		ID:    id,
		Batch: input.Batch,
		Year:  input.Year,
		Month: input.Month,
	}
}

func generateID() uint {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Int63n(9999999999-1000000000) + 1000000000
	return uint(randomNumber)
}

var mu sync.Mutex

func generateIDUSer() uint {
	mu.Lock()
	defer mu.Unlock()

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Int63n(9999999999-1000000000) + 1000000000
	return uint(randomNumber)
}
