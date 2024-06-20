package handler

import (
	"jastip-jakarta/features/user"
	"net/http"

	"jastip-jakarta/utils/responses"
	"jastip-jakarta/utils/middlewares"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}


func (handler *UserHandler) RegisterUser(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}

	userCore := RequestToUser(newUser)
	errInsert := handler.userService.Create(userCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Pendaftaran Berhasil", nil))
}


func (handler *UserHandler) GetUser(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)

	result, errSelect := handler.userService.GetById(userIdLogin)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}

	var userResult = UserToResponse(result)
	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil Get Data Profile", userResult))
}

func (handler *UserHandler) UpdateUser(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)

	var userData = UserUpdateRequest{}
	errBind := c.Bind(&userData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	fileHeader, _ := c.FormFile("photo_profile")

	userCore := UpdateRequestToUser(userData)
	errUpdate := handler.userService.Update(userIdLogin, userCore, fileHeader)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil edit profile", nil))
}

func (handler *UserHandler) Login(c echo.Context) error {
	var reqData = LoginRequest{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}
	result, token, err := handler.userService.Login(reqData.EmailOrPhone, reqData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}
	responseData := map[string]any{
		"token": token,
		"nama":  result.Name,
	}
	return c.JSON(http.StatusOK, responses.WebResponse("Login berhasil", responseData))
}
