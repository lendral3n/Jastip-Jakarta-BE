package service

import (
	"errors"
	"jastip-jakarta/features/user"
	"jastip-jakarta/utils/encrypts"
	"jastip-jakarta/utils/middlewares"
)

type userService struct {
	userData    user.UserDataInterface
	hashService encrypts.HashInterface
}

// dependency injection
func New(repo user.UserDataInterface, hash encrypts.HashInterface) user.UserServiceInterface {
	return &userService{
		userData:    repo,
		hashService: hash,
	}
}

// Create implements user.UserServiceInterface.
func (u *userService) Create(input user.User) error {
	if input.Name == "" {
		return errors.New("Nama tidak boleh kosong")
	}
	if input.Email == "" {
		return errors.New("Email tidak boleh kosong")
	}
	if input.PhoneNumber == 0 {
		return errors.New("Nomor Telephone tidak boleh kosong")
	}

	if input.Password != "" {
		hashedPass, errHash := u.hashService.HashPassword(input.Password)
		if errHash != nil {
			return errors.New("Error hash password.")
		}
		input.Password = hashedPass
	}

	err := u.userData.Insert(input)
	return err
}

// GetById implements user.UserServiceInterface.
func (u *userService) GetById(userIdLogin int) (*user.User, error) {
	userData, err := u.userData.SelectById(userIdLogin)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

// Login implements user.UserServiceInterface.
func (u *userService) Login(phoneOrEmail string, password string) (data *user.User, token string, err error) {
	// Validasi jika email atau password kosong
	if phoneOrEmail == "" {
		return nil, "", errors.New("Email atau nomor telepon tidak boleh kosong")
	}
	if password == "" {
		return nil, "", errors.New("Password tidak boleh kosong")
	}

	data, err = u.userData.Login(phoneOrEmail, password)
	if err != nil {
		return nil, "", err
	}

	isValid := u.hashService.CheckPasswordHash(data.Password, password)
	if !isValid {
		return nil, "", errors.New("Sandi Salah")
	}

	token, errJwt := middlewares.CreateToken(int(data.ID))
	if errJwt != nil {
		return nil, "", errJwt
	}

	return data, token, err
}

// Update implements user.UserServiceInterface.
func (u *userService) Update(userIdLogin int, input user.User) error {
	// Hash password baru jika diubah
	if input.Password != "" {
		hashedPass, errHash := u.hashService.HashPassword(input.Password)
		if errHash != nil {
			return errors.New("Error hash password.")
		}
		input.Password = hashedPass
	}

	err := u.userData.Update(userIdLogin, input)
	return err
}