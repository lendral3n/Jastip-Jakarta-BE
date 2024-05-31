package data

import (
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
	WhatsAppNumber int
	RegionCode     string
	AdminOrder     AdminOrder
}

type AdminOrder struct {
	gorm.Model
	UserOrderID           uint
	Status                string
	WeightItem            float64
	DeliveryBatch         string
	PackageWrappedPhoto   string
	PackageReceivedPhoto  string
	EstimatedDeliveryTime *time.Time
}

func AdminOrderStatusToModel(updateStatus order.AdminOrder) AdminOrder {
	return AdminOrder{
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
		WhatsAppNumber: input.WhatsAppNumber,
		RegionCode:     input.RegionCode,
	}
}

func (o UserOrder) ModelToUserOrder() order.UserOrder {
	return order.UserOrder{
		ID:             o.ID,
		UserID:         o.UserID,
		ItemName:       o.ItemName,
		TrackingNumber: o.TrackingNumber,
		OnlineStore:    o.OnlineStore,
		WhatsAppNumber: o.WhatsAppNumber,
		RegionCode:     o.RegionCode,
	}
}

func AdminOrderToModel(input order.AdminOrder) AdminOrder {
	return AdminOrder{
		UserOrderID:           input.UserOrderID,
		Status:                input.Status,
		WeightItem:            input.WeightItem,
		DeliveryBatch:         input.DeliveryBatch,
		PackageWrappedPhoto:   input.PackageWrappedPhoto,
		PackageReceivedPhoto:  input.PackageReceivedPhoto,
		EstimatedDeliveryTime: input.EstimatedDeliveryTime,
	}
}

func (o AdminOrder) ModelToAdminOrder() order.AdminOrder {
	return order.AdminOrder{
		ID:                    o.ID,
		UserOrderID:           o.UserOrderID,
		Status:                o.Status,
		WeightItem:            o.WeightItem,
		DeliveryBatch:         o.DeliveryBatch,
		PackageWrappedPhoto:   o.PackageWrappedPhoto,
		PackageReceivedPhoto:  o.PackageReceivedPhoto,
		EstimatedDeliveryTime: o.EstimatedDeliveryTime,
	}
}
