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
	DeliveryBatch        string `json:"delivery_batch,omitempty"`
	Code                 string `json:"code"`
	Region               string `json:"region"`
	Estimasi             string `json:"estimasi"`
	TotalOrder           int    `json:"total_order"`
	TotalWeight          int    `json:"total_weight"`
	TotalPrice           int    `json:"total_price"`
	PackageWrappedPhoto  string `json:"package_wrapped_photo"`
	PackageReceivedPhoto string `json:"package_received_photo"`
	Orders               []UserOrderProcessResponse `json:"orders"`
}

type GroupedAdminOrderResponse struct {
	DeliveryBatch  string        `json:"delivery_batch"`
	Code           string        `json:"code"`
	Region         string        `json:"region"`
	Estimasi       string        `json:"estimasi"`
	TotalOrder     int           `json:"total_order"`
	TotalWeight    int           `json:"total_weight"`
	TotalPrice     int           `json:"total_price"`
	CustomerJastip []Customer    `json:"customer_jastip"`
	Orders         []UserOrderProcessResponse `json:"orders"`
}

type UserOrderProcessResponse struct {
	ID                   uint   `json:"order_id"`
	Name                 string `json:"name"`
	ItemName             string `json:"item_name"`
	Status               string `json:"status"`
	TrackingNumberJastip string `json:"tracking_number_jastip"`
	TrackingNumber       string `json:"tracking_number"`
	OnlineStore          string `json:"online_store"`
	WeightItem           int    `json:"weight_item"`
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

type GetCustomerResponse struct {
	DeliveryBatch  string     `json:"delivery_batch"`
	Code           string     `json:"code"`
	Region         string     `json:"region"`
	Estimasi       string     `json:"estimasi"`
	CustomerJastip []Customer `json:"customer_jastip"`
}

type Customer struct {
	Name string `json:"name"`
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

func CoreToGetCustomerResponse(data []order.UserOrder, batch string, code string) GetCustomerResponse {
	var customers []Customer
	for _, userOrder := range data {
		customers = append(customers, Customer{Name: userOrder.User.Name})
	}

	region := ""
	if len(data) > 0 {
		region = data[0].Region.Region
	}

	estimasi := ""
	if len(data) > 0 && data[0].OrderDetails.EstimatedDeliveryTime != nil {
		estimasi = time.FormatDateToIndonesian(*data[0].OrderDetails.EstimatedDeliveryTime)
	}

	return GetCustomerResponse{
		DeliveryBatch:  batch,
		Code:           code,
		Region:         region,
		Estimasi:       estimasi,
		CustomerJastip: customers,
	}
}

// total berat, berat item + berat item
// total barang, item + item
// total harga, berat item x kode wilayah

func CoreToGroupedOrderResponse(data []order.UserOrder) GroupedOrderResponse {
	var totalWeight, totalPrice int
	var orders []UserOrderProcessResponse

	if len(data) == 0 {
		return GroupedOrderResponse{}
	}

	batch := data[0].OrderDetails.DeliveryBatchID
	code := data[0].Region.ID

	for _, userOrder := range data {
		orders = append(orders, CoreToUserOrderProcessResponse(userOrder))
		totalWeight += hitungTotalBerat([]order.UserOrder{userOrder})
		totalPrice += hitungTotalHarga([]order.UserOrder{userOrder})
	}

	estimasi := ""
	if data[0].OrderDetails.EstimatedDeliveryTime != nil {
		estimasi = time.FormatDateToIndonesian(*data[0].OrderDetails.EstimatedDeliveryTime)
	}

	return GroupedOrderResponse{
		DeliveryBatch: *batch,
		Code:          code,
		Region:        data[0].Region.Region,
		Estimasi:      estimasi,
		TotalOrder:    hitungTotalPesanan(data),
		TotalWeight:   totalWeight,
		TotalPrice:    totalPrice,
		// PackageReceivedPhoto: ,
		// PackageWrappedPhoto: ,
		Orders:        orders,
	}
}

func CoreToUserOrderProcessResponse(data order.UserOrder) UserOrderProcessResponse {
	return UserOrderProcessResponse{
		ID:                   data.ID,
		Name:                 data.User.Name,
		ItemName:             data.ItemName,
		Status:               data.OrderDetails.Status,
		TrackingNumberJastip: data.OrderDetails.TrackingNumberJastip,
		TrackingNumber:       data.TrackingNumber,
		OnlineStore:          data.OnlineStore,
		WeightItem:           int(data.OrderDetails.WeightItem),
	}
}

func hitungTotalPesanan(data []order.UserOrder) int {
	return len(data)
}

func hitungTotalBerat(data []order.UserOrder) int {
	totalBerat := 0
	for _, pesanan := range data {
		totalBerat += int(pesanan.OrderDetails.WeightItem)
	}
	return totalBerat
}

func hitungTotalHarga(data []order.UserOrder) int {
	totalHarga := 0
	for _, pesanan := range data {
		totalHarga += int(pesanan.OrderDetails.WeightItem) * dapatkanHargaPerBerat(pesanan.Region.Price)
	}
	return totalHarga
}

func dapatkanHargaPerBerat(harga int) int {
	return harga
}

func CoreToGroupedAdminOrderResponse(data []order.UserOrder, batch string, code string) GroupedAdminOrderResponse {
	var totalWeight, totalPrice int
	var orders []UserOrderProcessResponse
	var customers []Customer

	if len(data) == 0 {
		return GroupedAdminOrderResponse{}
	}

	region := data[0].Region.Region

	for _, userOrder := range data {
		orders = append(orders, CoreToUserOrderProcessResponse(userOrder))
		totalWeight += int(userOrder.OrderDetails.WeightItem)
		totalPrice += int(userOrder.OrderDetails.WeightItem) * userOrder.Region.Price
		customers = append(customers, Customer{Name: userOrder.User.Name})
	}

	estimasi := ""
	if data[0].OrderDetails.EstimatedDeliveryTime != nil {
		estimasi = time.FormatDateToIndonesian(*data[0].OrderDetails.EstimatedDeliveryTime)
	}

	return GroupedAdminOrderResponse{
		DeliveryBatch:  batch,
		Code:           code,
		Region:         region,
		Estimasi:       estimasi,
		TotalOrder:     len(data),
		TotalWeight:    totalWeight,
		TotalPrice:     totalPrice,
		CustomerJastip: customers,
		Orders:         orders,
	}
}
