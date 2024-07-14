package data

import (
	ad "jastip-jakarta/features/admin/data"
	"jastip-jakarta/features/order"
	ud "jastip-jakarta/features/user/data"

	"time"

	"gorm.io/gorm"
)

type UserOrder struct {
	ID uint `gorm:"primaryKey" json:"id"`
	gorm.Model
	UserID         uint
	ItemName       string
	TrackingNumber string
	OnlineStore    string
	WhatsappNumber int
	RegionCodeID   string
	User           ud.User       `gorm:"foreignKey:UserID"`
	Region         ad.RegionCode `gorm:"foreignKey:RegionCodeID"`
	OrderDetail    OrderDetail
	PhotoOrders    []PhotoOrder `gorm:"many2many:photo_order_user_orders;"`
}

type OrderDetail struct {
	gorm.Model
	UserOrderID           uint
	AdminID               *uint `gorm:"default:null"`
	Status                string
	WeightItem            int
	TrackingNumberJastip  string
	DeliveryBatchID       *string `gorm:"default:null"`
	EstimatedDeliveryTime *time.Time
	Admin                 ad.Admin         `gorm:"foreignKey:AdminID"`
	DeliveryBatch         ad.DeliveryBatch `gorm:"foreignKey:DeliveryBatchID"`
}

type PhotoOrder struct {
	gorm.Model
	DeliveryBatchID string
	PhotoPacked     string
	PhotoReceived   string
	DeliveryBatch   ad.DeliveryBatch `gorm:"foreignKey:DeliveryBatchID"`
	UserOrders      []UserOrder      `gorm:"many2many:photo_order_user_orders;"`
}

type PhotoOrderUserOrder struct {
	PhotoOrderID uint
	UserOrderID  uint
}

func OrderDetailStatusToModel(updateStatus order.OrderDetail) OrderDetail {
	return OrderDetail{
		Status: updateStatus.Status,
	}
}

func PhotoOrderToModel(input order.PhotoOrder) PhotoOrder {
	userOrders := []UserOrder{}
	for _, userID := range input.UserOrderIDs {
		userOrders = append(userOrders, UserOrder{ID: userID})
	}

	return PhotoOrder{
		DeliveryBatchID: input.DeliveryBatchID,
		PhotoPacked:     input.PhotoPacked,
		PhotoReceived:   input.PhotoReceived,
		UserOrders:      userOrders,
	}
}

func (o PhotoOrder) ModelToPhotoOrder() order.PhotoOrder {
	userOrderIDs := []uint{}
	for _, userOrder := range o.UserOrders {
		userOrderIDs = append(userOrderIDs, userOrder.ID)
	}

	return order.PhotoOrder{
		ID:              o.ID,
		UserOrderIDs:    userOrderIDs,
		DeliveryBatchID: o.DeliveryBatchID,
		PhotoPacked:     o.PhotoPacked,
		PhotoReceived:   o.PhotoReceived,
		CreatedAt:       o.CreatedAt,
		UpdatedAt:       o.UpdatedAt,
	}
}

func UserOrderToModel(input order.UserOrder) UserOrder {
	return UserOrder{
		ID:             input.ID,
		UserID:         input.UserID,
		ItemName:       input.ItemName,
		TrackingNumber: input.TrackingNumber,
		OnlineStore:    input.OnlineStore,
		WhatsappNumber: input.WhatsAppNumber,
		RegionCodeID:   input.RegionCode,
	}
}

func (uo UserOrder) ModelToUserOrderWait() order.UserOrder {
	return order.UserOrder{
		ID:             uo.ID,
		UserID:         uo.UserID,
		ItemName:       uo.ItemName,
		TrackingNumber: uo.TrackingNumber,
		OnlineStore:    uo.OnlineStore,
		WhatsAppNumber: uo.WhatsappNumber,
		Region:         uo.Region.ModelToRegionCode(),
		User:           uo.User.ModelToUser(),
		OrderDetails:   uo.OrderDetail.ModelToOrderDetail(),
	}
}

func (o UserOrder) ModelToUserOrderWaits() *order.UserOrder {
	if o.OrderDetail.Status != "Menunggu Diterima" {
		return nil
	}
	return &order.UserOrder{
		ID:             o.ID,
		UserID:         o.UserID,
		ItemName:       o.ItemName,
		TrackingNumber: o.TrackingNumber,
		OnlineStore:    o.OnlineStore,
		WhatsAppNumber: o.WhatsappNumber,
		Region:         o.Region.ModelToRegionCode(),
		User:           o.User.ModelToUser(),
		OrderDetails:   o.OrderDetail.ModelToOrderDetail(),
	}
}

func OrderDetailToModel(input order.OrderDetail) OrderDetail {
	return OrderDetail{
		AdminID:               input.AdminID,
		Status:                input.Status,
		WeightItem:            input.WeightItem,
		DeliveryBatchID:       input.DeliveryBatchID,
		TrackingNumberJastip:  input.TrackingNumberJastip,
		EstimatedDeliveryTime: input.EstimatedDeliveryTime,
	}
}

func (o OrderDetail) ModelToOrderDetail() order.OrderDetail {
	return order.OrderDetail{
		ID:                    o.ID,
		UserOrderID:           o.UserOrderID,
		Status:                o.Status,
		WeightItem:            o.WeightItem,
		DeliveryBatchID:       o.DeliveryBatchID,
		EstimatedDeliveryTime: o.EstimatedDeliveryTime,
		TrackingNumberJastip:  o.TrackingNumberJastip,
	}
}
