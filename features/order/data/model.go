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
	User           ud.User `gorm:"foreignKey:UserID"`
	ItemName       string
	TrackingNumber string
	OnlineStore    string
	WhatsappNumber int
	RegionCodeID   string
	Region         ad.RegionCode `gorm:"foreignKey:RegionCodeID"`
	OrderDetail    OrderDetail
}

type OrderDetail struct {
	gorm.Model
	UserOrderID           uint
	AdminID               *uint    `gorm:"default:null"`
	Admin                 ad.Admin `gorm:"foreignKey:AdminID"`
	Status                string
	WeightItem            float64
	PackageWrappedPhoto   string
	PackageReceivedPhoto  string
	TrackingNumberJastip  string
	DeliveryBatchID       string
	DeliveryBatch         ad.DeliveryBatch `gorm:"foreignKey:DeliveryBatchID"`
	EstimatedDeliveryTime *time.Time
}

func OrderDetailStatusToModel(updateStatus order.OrderDetail) OrderDetail {
	return OrderDetail{
		Status: updateStatus.Status,
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

func (o UserOrder) ModelToUserOrderWait() order.UserOrder {
	return order.UserOrder{
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
		UserOrderID:           input.UserOrderID,
		AdminID:               input.AdminID,
		Status:                input.Status,
		WeightItem:            input.WeightItem,
		DeliveryBatchID:       input.DeliveryBatchID,
		PackageWrappedPhoto:   input.PackageWrappedPhoto,
		PackageReceivedPhoto:  input.PackageReceivedPhoto,
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
		PackageWrappedPhoto:   o.PackageWrappedPhoto,
		PackageReceivedPhoto:  o.PackageReceivedPhoto,
		EstimatedDeliveryTime: o.EstimatedDeliveryTime,
		TrackingNumberJastip:  o.TrackingNumberJastip,
	}
}