package order

import (
	ad "jastip-jakarta/features/admin"
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
	ID                    uint
	UserOrderID           uint
	AdminID               uint
	Status                string
	WeightItem            float64
	DeliveryBatch         string
	PackageWrappedPhoto   string
	PackageReceivedPhoto  string
	TrackingNumberJastip  string
	EstimatedDeliveryTime *time.Time
	Admin                 ad.Admin
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

// interface untuk Data Layer
type OrderDataInterface interface {
	InsertUserOrder(userIdLogin int, inputOrder UserOrder) error
	PutUserOrder(userIdLogin int, userOrderId uint, inputOrder UserOrder) error
	CheckOrderStatus(userOrderId uint) (string, error)
	SelectUserOrderWait(userIdLogin int) ([]UserOrder, error)
	SelectUserOrderProcess(userIdLogin int) ([]UserOrder, error)
	SelectById(IdOrder uint) (*UserOrder, error)
	InsertAdminOrder(adminIdLogin int, inputOrder AdminOrder) error
}

// interface untuk Service Layer
type OrderServiceInterface interface {
	CreateUserOrder(userIdLogin int, inputOrder UserOrder) error
	UpdateUserOrder(userIdLogin int, userOrderId uint, inputOrder UserOrder) error
	GetUserOrderWait(userIdLogin int) ([]UserOrder, error)
	GetUserOrderProcess(userIdLogin int) ([]UserOrder, error)
	GetById(IdOrder uint) (*UserOrder, error)
	CreateAdminOrder(adminIdLogin int, userOrderId uint, inputOrder AdminOrder) error
}
