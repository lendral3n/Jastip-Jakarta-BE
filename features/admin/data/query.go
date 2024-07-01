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
				return nil, errors.New("email tidak terdaftar")
			}
			return nil, err
		}
	} else {
		// Cari admin dengan Nomor Telepon
		phone, convErr := strconv.Atoi(phoneOrEmail)
		if convErr != nil {
			return nil, errors.New("format nomor telepon tidak valid")
		}
		err = u.db.Where("phone_number = ?", phone).First(&adminDataGorm).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("nomor telepon tidak terdaftar")
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

// SelectAllRegionCode implements admin.AdminDataInterface.
func (u *adminQuery) SelectAllRegionCode() ([]admin.RegionCode, error) {
	var regionCodes []RegionCode

	err := u.db.Find(&regionCodes).Error
	if err != nil {
		return nil, err
	}

	var responseRegionCodes []admin.RegionCode
	for _, rc := range regionCodes {
		responseRegionCodes = append(responseRegionCodes, rc.ModelToRegionCode())
	}

	return responseRegionCodes, nil
}

// SelectByIdRegion implements admin.AdminDataInterface.
func (a *adminQuery) SelectByIdRegion(IdRegion string) (*admin.RegionCode, error) {
	var regionCode admin.RegionCode
	err := a.db.Where("id = ?", IdRegion).First(&regionCode).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("kode wilayah tidak ditemukan")
		}
		return nil, err
	}
	return &regionCode, nil
}

// InsertBatchDelivery implements admin.AdminDataInterface.
func (u *adminQuery) InsertBatchDelivery(adminIdLogin int, input admin.DeliveryBatch) error {
	dataGormBatch := DeliveryBatchToModel(input)
	dataGormBatch.AdminID = uint(adminIdLogin)

	tx := u.db.Create(&dataGormBatch)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectAllBatchDelivery implements admin.AdminDataInterface.
func (u *adminQuery) SelectAllBatchDelivery() ([]admin.DeliveryBatch, error) {
	var batches []DeliveryBatch

	err := u.db.Find(&batches).Error
	if err != nil {
		return nil, err
	}

	var responseBatches []admin.DeliveryBatch
	for _, batch := range batches {
		responseBatches = append(responseBatches, batch.ModelToDeliveryBatch())
	}
	return responseBatches, nil
}
// SelectDeliveryBatch implements admin.AdminDataInterface.
func (u *adminQuery) SelectDeliveryBatch(batchID string) (*admin.DeliveryBatch, error) {
	var deliveryBatch admin.DeliveryBatch
	err := u.db.Where("id = ?", batchID).First(&deliveryBatch).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("batch pengiriman tidak ditemukan")
		}
		return nil, err
	}
	return &deliveryBatch, nil
}