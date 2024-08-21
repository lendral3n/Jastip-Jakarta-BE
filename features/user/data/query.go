package data

import (
	"errors"
	"jastip-jakarta/features/user"
	"jastip-jakarta/utils/cloudinary"
	"mime/multipart"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type userQuery struct {
	db  *gorm.DB
	cld cloudinary.CloudinaryUploaderInterface
}

func New(db *gorm.DB, cloudinaryUploader cloudinary.CloudinaryUploaderInterface) user.UserDataInterface {
	return &userQuery{
		db:  db,
		cld: cloudinaryUploader,
	}
}

// Insert implements user.UserDataInterface.
func (u *userQuery) Insert(input user.User) error {
	// Cek apakah email sudah ada
	var emailCheck User
	emailResult := u.db.Where("email = ?", input.Email).First(&emailCheck)
	if emailResult.RowsAffected > 0 {
		return errors.New("email sudah terdaftar")
	}

	// Cek apakah nama sudah ada
	var nameCheck User
	nameResult := u.db.Where("name = ?", input.Name).First(&nameCheck)
	if nameResult.RowsAffected > 0 {
		return errors.New("nama sudah terdaftar")
	}

	// Cek apakah nomor telepon sudah ada
	var phoneCheck User
	phoneResult := u.db.Where("phone_number = ?", input.PhoneNumber).First(&phoneCheck)
	if phoneResult.RowsAffected > 0 {
		return errors.New("nomor telepon sudah terdaftar")
	}

	// Jika tidak ada yang sama, lanjutkan dengan pembuatan akun baru
	dataGorm := UserToModel(input)
	tx := u.db.Create(&dataGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Login implements user.UserDataInterface.
func (u *userQuery) Login(phoneOrEmail, password string) (data *user.User, err error) {
	var userDataGorm User

	// Cek apakah input adalah email atau nomor telepon
	if strings.Contains(phoneOrEmail, "@") {
		// Cari user dengan Email
		err = u.db.Where("email = ?", phoneOrEmail).First(&userDataGorm).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("email tidak terdaftar")
			}
			return nil, err
		}
	} else {
		// Cari user dengan Nomor Telepon
		phone, convErr := strconv.Atoi(phoneOrEmail)
		if convErr != nil {
			return nil, errors.New("format nomor telepon tidak valid")
		}
		err = u.db.Where("phone_number = ?", phone).First(&userDataGorm).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("nomor telepon tidak terdaftar")
			}
			return nil, err
		}
	}

	userData := userDataGorm.ModelToUser()
	return &userData, nil
}

// SelectById implements user.UserDataInterface.
func (u *userQuery) SelectById(userIdLogin int) (*user.User, error) {
	var userDataGorm User
	err := u.db.Where("id = ?", userIdLogin).First(&userDataGorm).Error
	if err != nil {
		return nil, err
	}
	userData := userDataGorm.ModelToUser()
	return &userData, nil
}

// Update implements user.UserDataInterface.
func (u *userQuery) Update(userIdLogin int, input user.User, photo *multipart.FileHeader) error {
	// Cek apakah email sudah ada
	var emailCheck User
	emailResult := u.db.Where("email = ?", input.Email).First(&emailCheck)
	if emailResult.RowsAffected > 0 {
		return errors.New("email sudah terdaftar")
	}

	// Cek apakah nama sudah ada
	var nameCheck User
	nameResult := u.db.Where("name = ?", input.Name).First(&nameCheck)
	if nameResult.RowsAffected > 0 {
		return errors.New("nama sudah terdaftar")
	}

	// Cek apakah nomor telepon sudah ada
	var phoneCheck User
	phoneResult := u.db.Where("phone_number = ?", input.PhoneNumber).First(&phoneCheck)
	if phoneResult.RowsAffected > 0 {
		return errors.New("nomor telepon sudah terdaftar")
	}

	dataGorm := UserToModel(input)

	// Cek apakah ada file foto yang diupload
	if photo != nil {
		imageURL, err := u.cld.UploadImage(photo)
		if err != nil {
			return err
		}
		dataGorm.PhotoProfile = imageURL
	}

	tx := u.db.Model(&User{}).Where("id = ?", userIdLogin).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectByName finds a user by name
func (u *userQuery) SelectByName(name string) (*user.User, error) {
	var userDataGorm User
	err := u.db.Where("name = ?", name).First(&userDataGorm).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	userData := userDataGorm.ModelToUser()
	return &userData, nil
}

// SelectByNameOrEmail finds a user by name or email
func (u *userQuery) SelectByNameOrEmail(query string) ([]user.User, error) {
	var usersDataGorm []User
	err := u.db.Where("name LIKE ? OR email LIKE ?", "%"+query+"%", "%"+query+"%").Find(&usersDataGorm).Error
	if err != nil {
		return nil, err
	}

	// Convert the list of User models to the user.User domain models
	usersData := make([]user.User, len(usersDataGorm))
	for i, userDataGorm := range usersDataGorm {
		usersData[i] = userDataGorm.ModelToUser()
	}

	return usersData, nil
}

func (u *userQuery) UpdateUserByName(name string, input user.User) error {
	// Find user by name
	var existingUser User
	err := u.db.Where("name = ?", name).First(&existingUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	// Prepare the data for update
	dataGorm := UserToModel(input)

	// Update the user data
	tx := u.db.Model(&existingUser).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// SelectAllUser implements user.UserDataInterface.
func (u *userQuery) SelectAllUser() ([]user.User, error) {
	var usersDataGorm []User

	// Mengambil semua data pengguna dari database
	err := u.db.Find(&usersDataGorm).Error
	if err != nil {
		return nil, err
	}

	// Mengkonversi list User model ke list user.User domain models
	usersData := make([]user.User, len(usersDataGorm))
	for i, userDataGorm := range usersDataGorm {
		usersData[i] = userDataGorm.ModelToUser()
	}

	return usersData, nil
}
