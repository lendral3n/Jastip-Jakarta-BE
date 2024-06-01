package handler

import (
	"jastip-jakarta/features/order"
)

type UserOrderWaitResponse struct {
	Status         string `json:"status"`
	Name           string `json:"name"`
	ItemName       string `json:"item_name"`
	TrackingNumber string `json:"tracking_number"`
	OnlineStore    string `json:"online_store"`
	RegionCode     string `json:"region_code"`
}

func CoreToResponseUserOrderWait(data *order.UserOrder) UserOrderWaitResponse {
	return UserOrderWaitResponse{
		ItemName:       data.ItemName,
		TrackingNumber: data.TrackingNumber,
		OnlineStore:    data.OnlineStore,
		RegionCode:     data.RegionCode,
		Name:           data.User.Name,
		Status:         data.AdminOrders.Status,
	}
}
