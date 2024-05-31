package handler

import (
	"jastip-jakarta/features/order"
	"math/rand"
	"time"
)

type UserOrderRequest struct {
	ID             uint
	ItemName       string `json:"item_name"`
	TrackingNumber string `json:"tracking_number"`
	OnlineStore    string `json:"online_store"`
	WhatsAppNumber int    `json:"whatsapp_number"`
	RegionCode     string `json:"region_code"`
}

type AdminOrderRequest struct {
	Status                string    `json:"status"`
	WeightItem            float64   `json:"weight_item"`
	DeliveryBatch         string    `json:"delivery_path"`
	PackageWrappedPhoto   string    `json:"package_wrapped_photo"`
	PackageReceivedPhoto  string    `json:"package_received_photo"`
	EstimatedDeliveryTime *time.Time `json:"estimated_delivery"`
}

type UpdateStatusRequest struct {
	Status string `json:"status"`
}

func RequestToUserOrder(input UserOrderRequest) order.UserOrder {
	return order.UserOrder{
		ID:             generateID(),
		ItemName:       input.ItemName,
		TrackingNumber: input.TrackingNumber,
		OnlineStore:    input.OnlineStore,
		WhatsAppNumber: input.WhatsAppNumber,
		RegionCode:     input.RegionCode,
	}
}

func RequestUpdateToUserOrder(input UserOrderRequest) order.UserOrder {
	return order.UserOrder{
		ItemName:       input.ItemName,
		TrackingNumber: input.TrackingNumber,
		OnlineStore:    input.OnlineStore,
		WhatsAppNumber: input.WhatsAppNumber,
		RegionCode:     input.RegionCode,
	}
}

func RequestToAdminOrder(input AdminOrderRequest) order.AdminOrder {
	return order.AdminOrder{
		Status:                input.Status,
		WeightItem:            input.WeightItem,
		DeliveryBatch:         input.DeliveryBatch,
		PackageWrappedPhoto:   input.PackageWrappedPhoto,
		PackageReceivedPhoto:  input.PackageReceivedPhoto,
		EstimatedDeliveryTime: input.EstimatedDeliveryTime,
	}
}

func generateID() uint {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Int63n(99999-10000) + 10000
	return uint(randomNumber)
}
