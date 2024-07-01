package data

import (
	"jastip-jakarta/features/admin"

	"gorm.io/gorm"
)

type Admin struct {
	ID uint `gorm:"primaryKey" json:"id"`
	gorm.Model
	Name         string
	Email        string
	Password     string
	PhoneNumber  int
	PhotoProfile string
	Role         string
	RegionCode   RegionCode `gorm:"foreignKey:AdminID"`
}

type RegionCode struct {
	ID string `gorm:"type:varchar(255);primaryKey" json:"id"`
	gorm.Model
	Region      string
	FullAddress string
	PhoneNumber int
	Price       int
	AdminID     uint
}

type DeliveryBatch struct {
	ID string `gorm:"type:varchar(255);primaryKey" json:"id"`
	gorm.Model
	Batch   int
	Year    int
	Month   int
	AdminID uint
	Admin   Admin `gorm:"foreignKey:AdminID"`
}

func AdminToModel(input admin.Admin) Admin {
	return Admin{
		ID:           input.ID,
		Name:         input.Name,
		Email:        input.Email,
		Password:     input.Password,
		PhoneNumber:  input.PhoneNumber,
		PhotoProfile: input.PhotoProfile,
		Role:         input.Role,
	}
}

func (u Admin) ModelToAdmin() admin.Admin {
	return admin.Admin{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		Password:     u.Password,
		PhoneNumber:  u.PhoneNumber,
		PhotoProfile: u.PhotoProfile,
		Role:         u.Role,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func RegionCodeToModel(input admin.RegionCode) RegionCode {
	return RegionCode{
		ID:          input.ID,
		Region:      input.Region,
		FullAddress: input.FullAddress,
		PhoneNumber: input.PhoneNumber,
		Price:       input.Price,
		AdminID:     input.AdminID,
	}
}

func (u RegionCode) ModelToRegionCode() admin.RegionCode {
	return admin.RegionCode{
		ID:          u.ID,
		Region:      u.Region,
		FullAddress: u.FullAddress,
		PhoneNumber: u.PhoneNumber,
		Price:       u.Price,
		AdminID:     u.AdminID,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

func DeliveryBatchToModel(input admin.DeliveryBatch) DeliveryBatch {
	return DeliveryBatch{
		ID:    input.ID,
		Batch: input.Batch,
		Year:  input.Year,
		Month: input.Month,
	}
}

func (u DeliveryBatch) ModelToDeliveryBatch() admin.DeliveryBatch {
	return admin.DeliveryBatch{
		ID:    u.ID,
		Batch: u.Batch,
		Year:  u.Year,
		Month: u.Month,
	}
}
