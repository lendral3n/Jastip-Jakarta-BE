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
}

// interface untuk Data Layer
type AdminDataInterface interface {
	Insert(input Admin) error
	Update(userIdLogin int, photo *multipart.FileHeader) error
	SelectById(userIdLogin int) (*Admin, error)
	Login(phoneOrEmail, password string) (data *Admin, err error)
}

// interface untuk Service Layer
type AdminServiceInterface interface {
	CreateSuper(input Admin) error
	Create(userIdLogin int, input Admin) error
	GetById(userIdLogin int) (*Admin, error)
	Update(userIdLogin int, photo *multipart.FileHeader) error
	Login(phoneOrEmail, password string) (data *Admin, token string, err error)
}
