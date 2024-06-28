package admin

import (
	"mime/multipart"
	"time"
)

type Admin struct {
	ID           uint
	Name         string
	Email        string
	Password     string
	PhoneNumber  int
	PhotoProfile string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	RegionCodes  RegionCode
}

type RegionCode struct {
	ID          string
	Region      string
	FullAddress  string
	PhoneNumber int
	Price       int
	AdminID     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// interface untuk Data Layer
type AdminDataInterface interface {
	Insert(input Admin) error
	Update(adminIdLogin int, photo *multipart.FileHeader) error
	SelectById(adminIdLogin int) (*Admin, error)
	Login(phoneOrEmail, password string) (data *Admin, err error)
	InsertRegionCode(input RegionCode) error
	SelectAllRegionCode() ([]RegionCode, error)
	SelectByIdRegion(IdRegion string) (*RegionCode, error)
}

// interface untuk Service Layer
type AdminServiceInterface interface {
	CreateSuper(input Admin) error
	Create(adminIdLogin int, input Admin) error
	GetById(adminIdLogin int) (*Admin, error)
	Update(adminIdLogin int, photo *multipart.FileHeader) error
	Login(phoneOrEmail, password string) (data *Admin, token string, err error)
	CreateRegionCode(adminIdLogin int, input RegionCode) error
	GetAllRegionCode() ([]RegionCode, error)
	GettByIdRegion(IdRegion string) (*RegionCode, error)
}
