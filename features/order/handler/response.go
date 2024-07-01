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
	ID                   uint   `json:"order_id"`
	Status               string `json:"status"`
	Name                 string `json:"name"`
	ItemName             string `json:"item_name"`
	TrackingNumber       string `json:"tracking_number"`
	TrackingNumberJastip string `json:"tracking_number_jastip"`
	OnlineStore          string `json:"online_store"`
	Code                 string `json:"code"`
	Region               string `json:"region"`
	FullAddress          string `json:"full_address"`
	WhatsappNumber       int    `json:"whatsapp_number"`
	WeightItem           int    `json:"weight_item"`
}

type DeliveryBatchWithRegionResponse struct {
	DeliveryBatchID string               `json:"delivery_batch"`
	RegionCodes     []RegionCodeResponse `json:"region_code"`
}

type RegionCodeResponse struct {
	Code   string `json:"code"`
	Region string `json:"region"`
}

func CoreToResponseUserOrderById(data order.UserOrder) OrderResponseById {
	return OrderResponseById{
		ID:                   data.ID,
		ItemName:             data.ItemName,
		TrackingNumber:       data.TrackingNumber,
		OnlineStore:          data.OnlineStore,
		Region:               data.Region.Region,
		Code:                 data.Region.ID,
		FullAddress:          data.Region.FullAddress,
		WhatsappNumber:       data.WhatsAppNumber,
		WeightItem:           int(data.OrderDetails.WeightItem),
		Name:                 data.User.Name,
		Status:               data.OrderDetails.Status,
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

		if groupedOrders[order.OrderDetails.DeliveryBatchID] == nil {
			groupedOrders[order.OrderDetails.DeliveryBatchID] = make(map[string][]UserOrderProcessResponse)
		}

		if groupedOrders[order.OrderDetails.DeliveryBatchID][regionCode] == nil {
			groupedOrders[order.OrderDetails.DeliveryBatchID][regionCode] = make([]UserOrderProcessResponse, 0)
		}

		// Menambahkan pesanan ke dalam grup berdasarkan region code
		groupedOrders[order.OrderDetails.DeliveryBatchID][regionCode] = append(
			groupedOrders[order.OrderDetails.DeliveryBatchID][regionCode],
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

func CoreToResponseDeliveryBatches(data []order.DeliveryBatchWithRegion) []DeliveryBatchWithRegionResponse {
	// Menggunakan map untuk memastikan setiap kombinasi code-region unik dalam setiap delivery_batch
	finalResult := make([]DeliveryBatchWithRegionResponse, 0)
	uniqueMap := make(map[string]map[string]bool) // Map untuk tracking kombinasi unik

	for _, item := range data {
		key := item.DeliveryBatchID
		if _, ok := uniqueMap[key]; !ok {
			uniqueMap[key] = make(map[string]bool)
		}

		// Cek apakah kombinasi code-region sudah ada dalam delivery_batch ini
		regionKey := item.RegionCode + "-" + item.Region
		if _, found := uniqueMap[key][regionKey]; !found {
			uniqueMap[key][regionKey] = true

			// Tambahkan ke response jika belum ada
			foundResponse := false
			for i := range finalResult {
				if finalResult[i].DeliveryBatchID == item.DeliveryBatchID {
					finalResult[i].RegionCodes = append(finalResult[i].RegionCodes, RegionCodeResponse{
						Code:   item.RegionCode,
						Region: item.Region,
					})
					foundResponse = true
					break
				}
			}
			if !foundResponse {
				finalResult = append(finalResult, DeliveryBatchWithRegionResponse{
					DeliveryBatchID: item.DeliveryBatchID,
					RegionCodes: []RegionCodeResponse{
						{
							Code:   item.RegionCode,
							Region: item.Region,
						},
					},
				})
			}
		}
	}

	return finalResult
}

// total berat, berat item + berat item
// total barang, item + item
// total harga, berat item x kode wilayah
