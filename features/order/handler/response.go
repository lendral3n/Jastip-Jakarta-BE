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
	Code           string `json:"code"`
	Region         string `json:"region"`
}

type GroupedOrderResponse struct {
	DeliveryBatch string                     `json:"delivery_batch,omitempty"`
	Code          string                     `json:"code"`
	Region        string                     `json:"region"`
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
	TrackingNumberJastip string `json:"tracking_number_jastip"`
	OnlineStore    string `json:"online_store"`
	Code           string `json:"code"`
	Region         string `json:"region"`
	FullAddress    string `json:"full_address"`
	WhatsappNumber int    `json:"whatsapp_number"`
	WeightItem     int    `json:"weight_item"`
}

func CoreToResponseUserOrderById(data order.UserOrder) OrderResponseById {
	return OrderResponseById{
		ID:             data.ID,
		ItemName:       data.ItemName,
		TrackingNumber: data.TrackingNumber,
		OnlineStore:    data.OnlineStore,
		Region:         data.Region.Region,
		Code:           data.Region.ID,
		FullAddress:    data.Region.FullAddress,
		WhatsappNumber: data.WhatsAppNumber,
		WeightItem:     int(data.OrderDetails.WeightItem),
		Name:           data.User.Name,
		Status:         data.OrderDetails.Status,
		TrackingNumberJastip: data.OrderDetails.TrackingNumberJastip,
	}
}

func CoreToResponseUserOrderWait(data order.UserOrder) UserOrderWaitResponse {
	return UserOrderWaitResponse{
		ID:             data.ID,
		ItemName:       data.ItemName,
		TrackingNumber: data.TrackingNumber,
		OnlineStore:    data.OnlineStore,
		Code:           data.Region.ID,
		Region:         data.Region.Region,
		Name:           data.User.Name,
		Status:         data.OrderDetails.Status,
	}
}

func CoreToResponseUserOrderProcess(data []order.UserOrder) []GroupedOrderResponse {
	groupedOrders := make(map[string]map[string][]UserOrderProcessResponse)

	// Mengelompokkan pesanan berdasarkan delivery batch dan region code
	for _, order := range data {
		regionCode := order.Region.ID // Menggunakan ID region sebagai region code

		if groupedOrders[order.OrderDetails.DeliveryBatch] == nil {
			groupedOrders[order.OrderDetails.DeliveryBatch] = make(map[string][]UserOrderProcessResponse)
		}

		if groupedOrders[order.OrderDetails.DeliveryBatch][regionCode] == nil {
			groupedOrders[order.OrderDetails.DeliveryBatch][regionCode] = make([]UserOrderProcessResponse, 0)
		}

		// Menambahkan pesanan ke dalam grup berdasarkan region code
		groupedOrders[order.OrderDetails.DeliveryBatch][regionCode] = append(
			groupedOrders[order.OrderDetails.DeliveryBatch][regionCode],
			UserOrderProcessResponse{
				ID:                   order.ID,
				Name:                 order.User.Name,
				ItemName:             order.ItemName,
				Status:               order.OrderDetails.Status,
				TrackingNumberJastip: order.OrderDetails.TrackingNumberJastip,
				TrackingNumber:       order.TrackingNumber,
				OnlineStore:          order.OnlineStore,
				WeightItem:           int(order.OrderDetails.WeightItem),
			},
		)
	}

	// Mengonversi map ke slice dari GroupedOrderResponse
	var response []GroupedOrderResponse
	for deliveryBatch, regions := range groupedOrders {
		for regionCode, orders := range regions {
			estimasi := ""
			if len(data) > 0 && data[0].OrderDetails.EstimatedDeliveryTime != nil {
				estimasi = time.FormatDateToIndonesian(*data[0].OrderDetails.EstimatedDeliveryTime)
			}
			totalOrder := len(orders)
			totalWeight := 0
			for _, order := range orders {
				totalWeight += order.WeightItem
			}
			// TODO: Menghitung totalPrice berdasarkan region code, perlu akses ke data harga per region

			// Ambil Region dari RegionCode terkait
			region := data[0].Region.Region

			response = append(response, GroupedOrderResponse{
				DeliveryBatch: deliveryBatch,
				Code:          regionCode, // Menggunakan regionCode sebagai Code
				Region:        region,     // Menggunakan Region dari RegionCode terkait
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
