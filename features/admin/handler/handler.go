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

func (handler *AdminHandler) GetRegionCode(c echo.Context) error {
	regionCodes, err := handler.adminService.GetAllRegionCode()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}

	var regionCodeResponses []RegionCodeResponse
	for _, regionCode := range regionCodes {
		regionCodeResponses = append(regionCodeResponses, CoreToResponseRegionCode(regionCode))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil mengambil kode region", regionCodeResponses))
}

func (handler *AdminHandler) GetRegionCodeById(c echo.Context) error {
	IdRegion := c.Param("code")

	regionCode, err := handler.adminService.GettByIdRegion(IdRegion)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}

	// Convert data to response format
	regionCodeResponse := CoreToResponseRegionCode(*regionCode)

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil mengambil kode region", regionCodeResponse))
}

func (handler *AdminHandler) CreateDeliveryBatch(c echo.Context) error {
	adminIdLogin := middlewares.ExtractTokenUserId(c)
	if adminIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
	}

	newBatch := DeliveryBatchRequest{}
	errBind := c.Bind(&newBatch)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}

	batchCore := RequestToDeliveryBatch(newBatch)
	errInsert := handler.adminService.CreateBatchDelivery(adminIdLogin, batchCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Membuat Batch Baru Berhasil", nil))
}

func (handler *AdminHandler) GetAllDeliveryBatch(c echo.Context) error {
	deliveryBatches, err := handler.adminService.GetAllBatchDelivery()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}

	var deliveryBatchResponse []DeliveryBatchResponse
	for _, deliveryBatch := range deliveryBatches {
		deliveryBatchResponse = append(deliveryBatchResponse, CoreToResponseDeliveryBatch(deliveryBatch))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil mengambil semua batch pengiriman", deliveryBatchResponse))
}

func (handler *AdminHandler) GetDeliveryBatchById(c echo.Context) error {
	batchID := c.Param("batch_id")

	deliveryBatch, err := handler.adminService.GetDeliveryBatch(batchID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}

	deliveryBatchResponse := CoreToResponseDeliveryBatch(*deliveryBatch)
	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil mengambil data batch pengiriman", deliveryBatchResponse))
}