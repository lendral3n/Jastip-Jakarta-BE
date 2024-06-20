package data

import (
	"errors"
	"jastip-jakarta/features/admin"
	"jastip-jakarta/utils/cloudinary"
	"mime/multipart"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type adminQuery struct {
	db  *gorm.DB
	cld cloudinary.CloudinaryUploaderInterface
}

func New(db *gorm.DB, cloudinaryUploader cloudinary.CloudinaryUploaderInterface) admin.AdminDataInterface {
	return &adminQuery{
		db:  db,
		cld: cloudinaryUploader,
	}
}

// Insert implements admin.AdminDataInterface.
func (u *adminQuery) Insert(input admin.Admin) error {
	dataGorm := AdminToModel(input)

	tx := u.db.Create(&dataGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Login implements admin.AdminDataInterface.
func (u *adminQuery) Login(phoneOrEmail, password string) (data *admin.Admin, err error) {
	var adminDataGorm Admin

	// Cek apakah input adalah email atau nomor telepon
	if strings.Contains(phoneOrEmail, "@") {
		// Cari admin dengan Email
		err = u.db.Where("email = ?", phoneOrEmail).First(&adminDataGorm).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("Email tidak terdaftar")
			}
			return nil, err
		}
	} else {
		// Cari admin dengan Nomor Telepon
		phone, convErr := strconv.Atoi(phoneOrEmail)
		if convErr != nil {
			return nil, errors.New("Format nomor telepon tidak valid")
		}
		err = u.db.Where("phone_number = ?", phone).First(&adminDataGorm).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("Nomor telepon tidak terdaftar")
			}
			return nil, err
		}
	}

	adminData := adminDataGorm.ModelToAdmin()
	return &adminData, nil
}

// SelectById implements admin.AdminDataInterface.
func (u *adminQuery) SelectById(adminIdLogin int) (*admin.Admin, error) {
	var adminDataGorm Admin
	err := u.db.Where("id = ?", adminIdLogin).First(&adminDataGorm).Error
	if err != nil {
		return nil, err
	}
	adminData := adminDataGorm.ModelToAdmin()
	return &adminData, nil
}

// Update implements admin.AdminDataInterface.
func (u *adminQuery) Update(adminIdLogin int, photo *multipart.FileHeader) error {
	imageURL, err := u.cld.UploadImage(photo)
	if err != nil {
		return err
	}

	dataGorm := &Admin{}
	dataGorm.PhotoProfile = imageURL

	tx := u.db.Model(&Admin{}).Where("id = ?", adminIdLogin).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// InsertRegionCode implements admin.AdminDataInterface.
func (u *adminQuery) InsertRegionCode(input admin.RegionCode) error {
	dataGorm := RegionCodeToModel(input)

	tx := u.db.Create(&dataGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

