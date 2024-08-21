package user

import (
	"mime/multipart"
	"time"
)

type User struct {
	ID       uint
	Name         string
	Email        string
	Password     string
	PhoneNumber  int
	PhotoProfile string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// interface untuk Data Layer
type UserDataInterface interface {
	Insert(input User) error
	Update(userIdLogin int, input User, photo *multipart.FileHeader) error
	SelectById(userIdLogin int) (*User, error)
	Login(phoneOrEmail, password string) (data *User, err error)
	SelectByNameOrEmail(query string) ([]User, error)
	UpdateUserByName(name string, input User) error
	SelectAllUser() ([]User, error)
}

// interface untuk Service Layer
type UserServiceInterface interface {
	Create(input User) error
	GetById(userIdLogin int) (*User, error)
	Update(userIdLogin int, input User, photo *multipart.FileHeader) error
	Login(phoneOrEmail, password string) (data *User, token string, err error)
}
