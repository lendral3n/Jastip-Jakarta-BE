package order

import (
	ad "jastip-jakarta/features/admin"
	ud "jastip-jakarta/features/user"
	"mime/multipart"
	"time"
)

type UserOrder struct {
	ID             uint
	UserID         uint
	ItemName       string
	TrackingNumber string
	OnlineStore    string
	WhatsAppNumber int
	RegionCode     string
	Region         ad.RegionCode
	User           ud.User
	CreatedAt      time.Time
	UpdatedAt      time.Time
	OrderDetails   OrderDetail
	PhotoOrders    PhotoOrder
}

type DeliveryBatchWithRegion struct {
	UserOrderID     uint
	RegionCode      string
	Region          string
	DeliveryBatchID string
	DeliveryBatch   ad.DeliveryBatch
}

type OrderDetail struct {
	ID                    uint
	UserOrderID           uint
	AdminID               *uint
	Status                string
	WeightItem            int
	DeliveryBatchID       *string
	TrackingNumberJastip  string
	EstimatedDeliveryTime *time.Time
	DeliveryBatch         ad.DeliveryBatch
	Admin                 ad.Admin
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type PhotoOrder struct {
	ID              uint
	DeliveryBatchID string
	PhotoPacked     string
	PhotoReceived   string
	RegionCodeID    string
	UserID          uint
	User            ud.User
	Region          ad.RegionCode
	DeliveryBatch   ad.DeliveryBatch
}

// interface untuk Data Layer
type OrderDataInterface interface {
	InsertUserOrder(userIdLogin int, inputOrder UserOrder) error
	PutUserOrder(userIdLogin int, userOrderId uint, inputOrder UserOrder) error
	CheckOrderStatus(userOrderId uint) (string, error)
	SelectUserOrderWait(userIdLogin int) ([]UserOrder, error)
	SelectUserOrderProcess(userIdLogin int) ([]UserOrder, error)
	SelectById(IdOrder uint) (*UserOrder, error)
	SearchUserOrder(userIdLogin int, itemName string) ([]UserOrder, error)
	InsertOrderDetail(adminIdLogin int, userOrderId uint, inputOrder OrderDetail) error
	SelectAllUserOrderWait() ([]UserOrder, error)
	FetchDeliveryBatchWithRegion() ([]DeliveryBatchWithRegion, error)
	SelectNameByUserOrder(code, batch string) ([]UserOrder, error)
	SelectOrderByUserOrderNameUser(code, batch, name string) ([]UserOrder, error)
	UpdateEstimationForOrders(code, batch string, estimation *time.Time) error
	UpdateOrderStatus(userOrderId uint, status string) error
	UploadFotoPacked(inputOrder PhotoOrder, photoPacked *multipart.FileHeader) error
	UploadFotoReceived(idFoto uint, photoReceived *multipart.FileHeader) error
	FetchOrdersByBatch(batch string) ([]UserOrder, error)
	GenerateCSVByBatch(batch string, filePath string) error
	GetFoto(batch, code string, userId int) (*PhotoOrder, error)
	SearchOrders(searchQuery string) ([]UserOrder, error)
}

// interface untuk Service Layer
type OrderServiceInterface interface {
	CreateUserOrder(userIdLogin int, inputOrder UserOrder) error
	UpdateUserOrder(userIdLogin int, userOrderId uint, inputOrder UserOrder) error
	GetUserOrderWait(userIdLogin int) ([]UserOrder, error)
	GetUserOrderProcess(userIdLogin int) ([]UserOrder, error)
	GetById(IdOrder uint) (*UserOrder, error)
	SearchUserOrder(userIdLogin int, itemName string) ([]UserOrder, error)
	CreateOrderDetail(adminIdLogin int, userOrderId uint, inputOrder OrderDetail) error
	GetAllUserOrderWait(adminIdLogin int) ([]UserOrder, error)
	GetDeliveryBatchWithRegion(adminIdLogin int) ([]DeliveryBatchWithRegion, error)
	GetNameByUserOrder(adminIdLogin int, code, batch string) ([]UserOrder, error)
	GetOrderByUserOrderNameUser(adminIdLogin int, code, batch, name string) ([]UserOrder, error)
	UpdateEstimationForOrders(adminIdLogin int, code, batch string, estimation *time.Time) error
	UpdateOrderStatus(adminIdLogin int, userOrderId uint, status string) error
	UploadFotoPacked(adminIdLogin int, inputOrder PhotoOrder, photoPacked *multipart.FileHeader) error
	UploadFotoReceived(adminIdLogin int, idFoto uint, photoReceived *multipart.FileHeader) error
	GenerateCSVByBatch(batch, filePath string) error
	GetFoto(batch, code string, userId int) (*PhotoOrder, error)
	SearchOrders(adminIdLogin int, searchQuery string) ([]UserOrder, error)
}
