package handler

import (
	"fmt"
	"jastip-jakarta/features/admin"
	"math/rand"
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

// var (
// 	batchNumberMap = make(map[string]int)
// 	mu             sync.Mutex
// )

func generateID() uint {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Int63n(9999999999-1000000000) + 1000000000
	return uint(randomNumber)
}

// func generateBatchID() string {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	now := time.Now()
// 	month := now.Month()
// 	year := now.Year()

// 	monthStr := fmt.Sprintf("%02d", month)
// 	yearStr := fmt.Sprintf("%04d", year)
// 	key := fmt.Sprintf("%s%s", monthStr, yearStr)

// 	if _, exists := batchNumberMap[key]; !exists {
// 		batchNumberMap[key] = 1
// 	} else {
// 		batchNumberMap[key]++
// 	}

// 	batchID := fmt.Sprintf("%s%sB%d", monthStr, yearStr, batchNumberMap[key])
// 	return batchID
// }

// bulan + tahun + B + angka Batch dimulai 1
