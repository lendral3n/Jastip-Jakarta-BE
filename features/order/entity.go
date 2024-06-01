package order

import (
	ud "jastip-jakarta/features/user"
	"time"
)

type UserOrder struct {
	// ID uint
	ID             uint
	UserID         uint
	ItemName       string
	TrackingNumber string
	OnlineStore    string
	WhatsAppNumber int
	RegionCode     string
	User           ud.User
	CreatedAt      time.Time
	UpdatedAt      time.Time
	AdminOrders    AdminOrder
}

type AdminOrder struct {
	ID uint
	UserOrderID uint
	// UserID                uint
	Status                string
	WeightItem            float64
	DeliveryBatch         string
	PackageWrappedPhoto   string
	PackageReceivedPhoto  string
	EstimatedDeliveryTime *time.Time
	// User                  ud.Core
	CreatedAt time.Time
	UpdatedAt time.Time
}

// interface untuk Data Layer
type OrderDataInterface interface {
	InsertUserOrder(userIdLogin int, inputOrder UserOrder) error
	PutUserOrder(userIdLogin int, userOrderId uint, inputOrder UserOrder) error
	CheckOrderStatus(userOrderId uint) (string, error)
	SelectUserOrderWait(userIdLogin int) ([]UserOrder, error)
}

// interface untuk Service Layer
type OrderServiceInterface interface {
	CreateOrder(userIdLogin int, inputOrder UserOrder) error
	UpdateUserOrder(userIdLogin int, userOrderId uint, inputOrder UserOrder) error
	SelectUserOrderWait(userIdLogin int) ([]UserOrder, error)
}
