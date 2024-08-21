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

	var adminResult = AdminToResponse(*result)
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

func (handler *AdminHandler) GetAdminJakarta(c echo.Context) error {
	adminIdLogin := middlewares.ExtractTokenUserId(c)
	if adminIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
	}

	role := "Jakarta"
	admin, err := handler.adminService.GetAdminsByRole(adminIdLogin, role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}

	var adminReponse []AdminResponse
	for _, admins := range admin {
		adminReponse = append(adminReponse, AdminToResponse(admins))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil mengambil admin jakarta", adminReponse))
}

func (handler *AdminHandler) GetAdminPerwakilan(c echo.Context) error {
	adminIdLogin := middlewares.ExtractTokenUserId(c)
	if adminIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
	}

	role := "Perwakilan"
	admin, err := handler.adminService.GetAdminsByRole(adminIdLogin, role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}

	var adminReponse []AdminResponse
	for _, admins := range admin {
		adminReponse = append(adminReponse, AdminToResponse(admins))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil mengambil admin perwakilan", adminReponse))
}

func (handler *AdminHandler) GetAllAdmin(c echo.Context) error {
	adminIdLogin := middlewares.ExtractTokenUserId(c)
	if adminIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
	}

	admin, err := handler.adminService.GetAllAdmins(adminIdLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}

	var adminReponse []AdminResponse
	for _, admins := range admin {
		adminReponse = append(adminReponse, AdminToResponse(admins))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil mengambil admin perwakilan", adminReponse))
}

func (handler *AdminHandler) SearchRegionCode(c echo.Context) error {
	adminIdLogin := middlewares.ExtractTokenUserId(c)
	if adminIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
	}

	query := c.QueryParam("code")

	regionCodes, err := handler.adminService.SearchRegionCode(adminIdLogin, query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error searching region code. "+err.Error(), nil))
	}

	var codeResponse []RegionCodeResponse
	for _, codes := range regionCodes {
		codeResponse = append(codeResponse, CoreToResponseRegionCode(codes))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil mencari kode wilayah", codeResponse))
}

func (handler *AdminHandler) UpdateRegionCode(c echo.Context) error {
	adminIdLogin := middlewares.ExtractTokenUserId(c)
	if adminIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
	}

    id := c.Param("code")
    if id == "" {
        return c.JSON(http.StatusBadRequest, responses.WebResponse("code tidak boleh kosong", nil))
    }

    updateRegion := RegionCodeRequest{}
	errBind := c.Bind(&updateRegion)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}

	regionCore := RequestToRegionCode(updateRegion)
    err := handler.adminService.UpdateRegionCode(adminIdLogin, id, regionCore)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.WebResponse("Failed to update region code. "+err.Error(), nil))
    }

    return c.JSON(http.StatusOK, responses.WebResponse("Region code updated successfully", nil))
}

func (handler *AdminHandler) SearchUser(c echo.Context) error {
    adminIdLogin := middlewares.ExtractTokenUserId(c)
    if adminIdLogin == 0 {
        return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
    }

    query := c.QueryParam("name")
    if query == "" {
        return c.JSON(http.StatusBadRequest, responses.WebResponse("Query parameter 'name or email' is required", nil))
    }

    users, err := handler.adminService.SearchUser(adminIdLogin, query)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.WebResponse("Error searching users. "+err.Error(), nil))
    }

    var userResponses []UserResponse
    for _, user := range users {
        userResponses = append(userResponses, UserToResponse(user))
    }

    return c.JSON(http.StatusOK, responses.WebResponse("Berhasil mencari pengguna", userResponses))
}

func (handler *AdminHandler) UpdateUserByName(c echo.Context) error {
    adminIdLogin := middlewares.ExtractTokenUserId(c)
    if adminIdLogin == 0 {
        return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
    }

    name := c.Param("name")
    if name == "" {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("name tidak boleh kosong", nil))
	}

    var updateUserReq UserRequest
    err := c.Bind(&updateUserReq)
    if err != nil {
        return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
    }

    userCore := RequestToUser(updateUserReq)
    
    err = handler.adminService.UpdateUserByName(adminIdLogin, name, userCore)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.WebResponse("Failed to update user. "+err.Error(), nil))
    }

    return c.JSON(http.StatusOK, responses.WebResponse("User updated successfully", nil))
}

func (handler *AdminHandler) CreateUser(c echo.Context) error {
	adminIdLogin := middlewares.ExtractTokenUserId(c)
    if adminIdLogin == 0 {
        return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
    }

	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}

	userCore := RequestRegisterToUser(newUser)
	errInsert := handler.adminService.CreateUser(adminIdLogin, userCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Pendaftaran Berhasil", nil))
}

func (handler *AdminHandler) GetAllUSer(c echo.Context) error {
	adminIdLogin := middlewares.ExtractTokenUserId(c)
    if adminIdLogin == 0 {
        return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
    }

	user, err := handler.adminService.GetAllUser(adminIdLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}

	var userResps []UserResponse
	for _, users := range user {
		userResps = append(userResps, UserToResponse(users))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil mengambil semua user", userResps))
}