package handler

import (
	"jastip-jakarta/features/order"
	"jastip-jakarta/utils/time"
)

type UserOrderWaitResponse struct {
	ID             uint   `json:"order_id"`
	Status         string `json:"status"`
	Name           string `json:"name"`
	ItemName       string `json:"item_name"`
	TrackingNumber string `json:"tracking_number"`
	OnlineStore    string `json:"online_store"`
	RegionCode     string `json:"region_code"`
}

type GroupedOrderResponse struct {
	DeliveryBatch string                     `json:"delivery_batch,omitempty"`
	RegionCode    string                     `json:"region_code"`
	Estimasi      string                     `json:"estimasi"`
	TotalOrder    int                        `json:"total_order"`
	TotalWeight   int                        `json:"total_weight"`
	TotalPrice    int                        `json:"total_price"`
	Orders        []UserOrderProcessResponse `json:"orders"`
}

type UserOrderProcessResponse struct {
	ID                   uint   `json:"order_id"`
	Name                 string `json:"name"`
	ItemName             string `json:"item_name"`
	Status               string `json:"status"`
	TrackingNumberJastip string `json:"tracking_number_jastip"`
	TrackingNumber       string `json:"tracking_number"`
	OnlineStore          string `json:"online_store"`
	WeightItem           int
}

type OrderResponseById struct {
	ID             uint   `json:"order_id"`
	Status         string `json:"status"`
	Name           string `json:"name"`
	ItemName       string `json:"item_name"`
	TrackingNumber string `json:"tracking_number"`
	OnlineStore    string `json:"online_store"`
	RegionCode     string `json:"region_code"`
}

func CoreToResponseUserOrderById(data order.UserOrder) OrderResponseById {
	return OrderResponseById{
		ID:             data.ID,
		ItemName:       data.ItemName,
		TrackingNumber: data.TrackingNumber,
		OnlineStore:    data.OnlineStore,
		RegionCode:     data.RegionCode,
		Name:           data.User.Name,
		Status:         data.AdminOrders.Status,
	}
}

func CoreToResponseUserOrderWait(data order.UserOrder) UserOrderWaitResponse {
	return UserOrderWaitResponse{
		ID:             data.ID,
		ItemName:       data.ItemName,
		TrackingNumber: data.TrackingNumber,
		OnlineStore:    data.OnlineStore,
		RegionCode:     data.RegionCode,
		Name:           data.User.Name,
		Status:         data.AdminOrders.Status,
	}
}

func CoreToResponseUserOrderProcess(data []order.UserOrder) []GroupedOrderResponse {
	groupedOrders := make(map[string]map[string][]UserOrderProcessResponse)

	// Mengelompokkan pesanan berdasarkan delivery batch dan region code
	for _, order := range data {
		if groupedOrders[order.AdminOrders.DeliveryBatch] == nil {
			groupedOrders[order.AdminOrders.DeliveryBatch] = make(map[string][]UserOrderProcessResponse)
		}
		groupedOrders[order.AdminOrders.DeliveryBatch][order.RegionCode] = append(
			groupedOrders[order.AdminOrders.DeliveryBatch][order.RegionCode],
			UserOrderProcessResponse{
				ID:                   order.ID,
				Name:                 order.User.Name,
				ItemName:             order.ItemName,
				Status:               order.AdminOrders.Status,
				TrackingNumberJastip: order.AdminOrders.TrackingNumberJastip,
				TrackingNumber:       order.TrackingNumber,
				OnlineStore:          order.OnlineStore,
				WeightItem:           int(order.AdminOrders.WeightItem),
			},
		)
	}

	// Mengonversi map ke slice dari GroupedOrderResponse
	var response []GroupedOrderResponse
	for deliveryBatch, regions := range groupedOrders {
		for regionCode, orders := range regions {
			estimasi := ""
			if len(data) > 0 && data[0].AdminOrders.EstimatedDeliveryTime != nil {
				estimasi = time.FormatDateToIndonesian(*data[0].AdminOrders.EstimatedDeliveryTime)
			}
			totalOrder := len(orders)
			totalWeight := 0
			for _, order := range orders {
				totalWeight += int(order.WeightItem)
			}
			// totalPrice := totalWeight * regionPriceMap[regionCode] // Pastikan Anda memiliki map harga per region code

			response = append(response, GroupedOrderResponse{
				DeliveryBatch: deliveryBatch,
				RegionCode:    regionCode,
				Estimasi:      estimasi,
				TotalOrder:    totalOrder,
				TotalWeight:   totalWeight,
				// TotalPrice:    totalPrice,
				Orders: orders,
			})
		}
	}

	return response
}

// total berat, berat item + berat item
// total barang, item + item
// total harga, berat item x kode wilayah
