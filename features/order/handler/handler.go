package handler

import (
	"jastip-jakarta/features/order"
	"jastip-jakarta/utils/middlewares"
	"jastip-jakarta/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderService order.OrderServiceInterface
}

func New(os order.OrderServiceInterface) *OrderHandler {
	return &OrderHandler{
		orderService: os,
	}
}

func (handler *OrderHandler) CreateUserOrder(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
	}

	newOrder := UserOrderRequest{}
	errBind := c.Bind(&newOrder)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data order not valid", nil))
	}

	orderCore := RequestToUserOrder(newOrder)
	errInsert := handler.orderService.CreateUserOrder(userIdLogin, orderCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil Membuat Orderan Jastip", nil))
}

func (handler *OrderHandler) UpdateUserOrder(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
	}

	userOrderId, errParse := strconv.ParseUint(c.Param("order_id"), 10, 32)
	if errParse != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("ID order tidak valid", nil))
	}

	updateOrder := UserOrderRequest{}
	errBind := c.Bind(&updateOrder)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data order not valid", nil))
	}

	orderCore := RequestUpdateToUserOrder(updateOrder)
	errUpdate := handler.orderService.UpdateUserOrder(userIdLogin, uint(userOrderId), orderCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Order berhasil diperbarui", nil))
}

func (handler *OrderHandler) GetUserOrderWait(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
	}

	userOrders, err := handler.orderService.GetUserOrderWait(userIdLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}

	var userOrderWaitResponses []UserOrderWaitResponse
	for _, userOrder := range userOrders {
		userOrderWaitResponses = append(userOrderWaitResponses, CoreToResponseUserOrderWait(userOrder))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil mendapatkan orderan yang menunggu dikirim", userOrderWaitResponses))
}

func (handler *OrderHandler) GetOrderById(c echo.Context) error {
	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("ID order tidak valid", nil))
	}

	result, errSelect := handler.orderService.GetById(uint(orderId))
	if errSelect != nil {
		return c.JSON(http.StatusNotFound, responses.WebResponse("Order tidak ditemukan", nil))
	}

	var orderResult = CoreToResponseUserOrderById(*result)
	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", orderResult))
}

func (handler *OrderHandler) CreateOrderDetail(c echo.Context) error {
	adminIdLogin := middlewares.ExtractTokenUserId(c)
	if adminIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
	}

	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("ID order tidak valid", nil))
	}

	newOrder := OrderDetailRequest{}
	errBind := c.Bind(&newOrder)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data order not valid", nil))
	}

	orderCore := RequestToOrderDetail(newOrder)
	errInsert := handler.orderService.CreateOrderDetail(adminIdLogin, uint(orderId), orderCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil Membuat Orderan Jastip", nil))
}

func (handler *OrderHandler) GetUserOrderProcess(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
	}

	userOrders, err := handler.orderService.GetUserOrderProcess(userIdLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}

	groupedResponses := CoreToResponseUserOrderProcess(userOrders)

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil mendapatkan orderan yang diproses", groupedResponses))
}

func (handler *OrderHandler) SearchUserOrder(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
	}

	itemName := c.QueryParam("item_name")
	if itemName == "" {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Nama barang tidak boleh kosong", nil))
	}

	userOrders, err := handler.orderService.SearchUserOrder(userIdLogin, itemName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}

	var userOrderResponses []UserOrderWaitResponse
	for _, userOrder := range userOrders {
		userOrderResponses = append(userOrderResponses, CoreToResponseUserOrderWait(userOrder))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil mencari orderan", userOrderResponses))
}

func (handler *OrderHandler) GetAllUserOrderWait(c echo.Context) error {
	adminIdLogin := middlewares.ExtractTokenUserId(c)
	if adminIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Silahkan login terlebih dahulu", nil))
	}

	userOrders, err := handler.orderService.GetAllUserOrderWait(adminIdLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(err.Error(), nil))
	}

	var userOrderResponses []UserOrderWaitResponse
	for _, userOrder := range userOrders {
		userOrderResponses = append(userOrderResponses, CoreToResponseUserOrderWait(userOrder))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Berhasil mendapatkan semua orderan yang menunggu diterima", userOrderResponses))
}
