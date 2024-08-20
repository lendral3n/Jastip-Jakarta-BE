package handler

import (
	"encoding/binary"
	"jastip-jakarta/features/order"
	"math/rand"
	"strconv"
	"time"
)

type UserOrderRequest struct {
	ID             uint
	ItemName       string `json:"item_name"`
	TrackingNumber string `json:"tracking_number"`
	OnlineStore    string `json:"online_store"`
	WhatsAppNumber int    `json:"whatsapp_number"`
	Code           string `json:"code"`
}

type OrderDetailRequest struct {
	Status        string  `json:"status"`
	WeightItem    float64 `json:"weight_item"`
	DeliveryBatch string  `json:"delivery_batch"`
}

type UploadFotoRequest struct {
	Batch  string `form:"batch"`
	Code   string `form:"code"`
	UserID uint   `form:"user_id"`
}

type UpdateStatusRequest struct {
	Status string `json:"status"`
}

type UpdateEstimationRequest struct {
	Estimation string `json:"estimation"`
}

type UpdateOrderByID struct {
	ItemName             string  `json:"item_name"`
	TrackingNumber       string  `json:"tracking_number"`
	OnlineStore          string  `json:"online_store"`
	WhatsAppNumber       int     `json:"whatsapp_number"`
	Code                 string  `json:"code"`
	WeightItem           float64 `json:"weight_item"`
	DeliveryBatch        string  `json:"delivery_batch"`
	TrackingNumberJastip string  `json:"tracking_number_jastip"`
}

func RequestToUserOrderUpdate(input UpdateOrderByID) order.UpdateOrderByID {
	return order.UpdateOrderByID{
		ItemName:             input.ItemName,
		TrackingNumber:       input.TrackingNumber,
		OnlineStore:          input.OnlineStore,
		WhatsAppNumber:       input.WhatsAppNumber,
		RegionCode:           input.Code,
		WeightItem:           input.WeightItem,
		DeliveryBatch:        input.DeliveryBatch,
		TrackingNumberJastip: input.TrackingNumberJastip,
	}
}

func RequestToUserOrder(input UserOrderRequest) order.UserOrder {
	return order.UserOrder{
		ID:             generateID(),
		ItemName:       input.ItemName,
		TrackingNumber: input.TrackingNumber,
		OnlineStore:    input.OnlineStore,
		WhatsAppNumber: input.WhatsAppNumber,
		RegionCode:     input.Code,
	}
}

func RequestToPhotoOrder(input UploadFotoRequest) order.PhotoOrder {
	return order.PhotoOrder{
		UserID:          input.UserID,
		DeliveryBatchID: input.Batch,
		RegionCodeID:    input.Code,
	}
}

func RequestUpdateToUserOrder(input UserOrderRequest) order.UserOrder {
	return order.UserOrder{
		ItemName:       input.ItemName,
		TrackingNumber: input.TrackingNumber,
		OnlineStore:    input.OnlineStore,
		WhatsAppNumber: input.WhatsAppNumber,
		RegionCode:     input.Code,
	}
}

func RequestToOrderDetail(input OrderDetailRequest, userOrder order.UserOrder) order.OrderDetail {
	deliveryBatch := input.DeliveryBatch
	return order.OrderDetail{
		Status:               input.Status,
		WeightItem:           input.WeightItem,
		TrackingNumberJastip: generateJastipResi(userOrder.WhatsAppNumber, userOrder.TrackingNumber),
		DeliveryBatchID:      &deliveryBatch,
	}
}

func generateID() uint {
	// Mendapatkan timestamp Unix saat ini dalam nanodetik
	timestamp := uint64(time.Now().UnixNano())

	// Menghasilkan angka acak 32-bit
	var randomBytes [4]byte
	if _, err := rand.Read(randomBytes[:]); err != nil {
		panic(err) // tangani error ini dengan tepat pada kode produksi
	}
	randomNumber := binary.BigEndian.Uint32(randomBytes[:])

	// Menggabungkan timestamp dan angka acak untuk membentuk ID unik
	uniqueID := (timestamp << 32) | uint64(randomNumber)
	return uint(uniqueID)
}

func ParseEstimationDate(estimation string) (*time.Time, error) {
	// Format tanggal dd/mm/yyyy
	layout := "02/01/2006"
	t, err := time.Parse(layout, estimation)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func generateJastipResi(wa int, resi string) string {
	// Konversi nomor WA ke string
	waStr := strconv.Itoa(wa)

	// Ambil 3 karakter terakhir dari resi
	var resiSuffix string
	if len(resi) >= 3 {
		resiSuffix = resi[len(resi)-3:]
	} else {
		resiSuffix = resi
	}

	// Gabungkan no_hp dan 3 karakter terakhir dari resi
	generatedResi := waStr + resiSuffix

	return generatedResi
}

// func RequestUpdateEstimasi(input UpdateEstimationRequest) (*time.Time, error) {
//     return ParseEstimationDate(input.Estimation)
// }
