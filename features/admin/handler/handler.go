package handler

import (
	"jastip-jakarta/features/admin"
	"net/http"

	"jastip-jakarta/utils/middlewares"
	"jastip-jakarta/utils/responses"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	adminService admin.AdminServiceInterface
}

func New(service admin.AdminServiceInterface) *AdminHandler {
	return &AdminHandler{
		adminService: service,
	}
}

func (handler *AdminHandler) RegisterAdminSuper(c echo.Context) error {
	newAdmin := AdminRequest{}
	errBind := c.Bind(&newAdmin)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}

	adminCore := RequestToAdmin(newAdmin)
	errInsert := handler.adminService.CreateSuper(adminCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Membuat Akun Admin Berhasil", nil))
}

func (handler *AdminHandler) RegisterAdmin(c echo.Context) error {
	adminIdLogin := middlewares.ExtractTokenUserId(c)
	newAdmin := AdminRequest{}
	errBind := c.Bind(&newAdmin)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}

	adminCore := RequestToAdmin(newAdmin)
	errInsert := handler.adminService.Create(adminIdLogin, adminCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Membuat Akun Admin Berhasil", nil))
}

func (handler *AdminHandler) GetAdmin(c echo.Context) error {
	adminIdLogin := middlewares.ExtractTokenUserId(c)

	result, errSelect := handler.adminService.GetById(adminIdLogin)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}

	var adminResult = AdminToResponse(result)
	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil Get Data Profile", adminResult))
}

func (handler *AdminHandler) UpdateAdmin(c echo.Context) error {
	adminIdLogin := middlewares.ExtractTokenUserId(c)

	fileHeader, err := c.FormFile("photo_profile")
	if err != nil && err != http.ErrMissingFile {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error retrieving the file", nil))
	}

	errUpdate := handler.adminService.Update(adminIdLogin, fileHeader)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil edit profile", nil))
}

func (handler *AdminHandler) Login(c echo.Context) error {
	var reqData = LoginRequest{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}
	result, token, err := handler.adminService.Login(reqData.EmailOrPhone, reqData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}
	responseData := map[string]any{
		"token": token,
		"nama":  result.Name,
		"role":  result.Role,
	}
	return c.JSON(http.StatusOK, responses.WebResponse("Login berhasil", responseData))
}

func (handler *AdminHandler) CreateRegionCode(c echo.Context) error {
	adminIdLogin := middlewares.ExtractTokenUserId(c)
	newRegion := RegionCodeRequest{}
	errBind := c.Bind(&newRegion)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}

	regionCore := RequestToRegionCode(newRegion)
	errInsert := handler.adminService.CreateRegionCode(adminIdLogin, regionCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Membuat Kode WIlayah Berhasil", nil))
}
