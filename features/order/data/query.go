package data

import (
	"jastip-jakarta/features/order"
	"log"

	"gorm.io/gorm"
)

type orderQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) order.OrderDataInterface {
	return &orderQuery{
		db: db,
	}
}

// InsertUserOrder implements order.OrderDataInterface.
func (o *orderQuery) InsertUserOrder(userIdLogin int, inputOrder order.UserOrder) error {
	newOrder := UserOrderToModel(inputOrder)
	newOrder.UserID = uint(userIdLogin)

	result := o.db.Create(&newOrder)
	if result.Error != nil {
		return result.Error
	}

	detailOrder := OrderDetail{
		UserOrderID: newOrder.ID,
		Status:      "Menunggu Diterima",
		AdminID: nil,
	}

	result = o.db.Create(&detailOrder)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// PutUserOrder implements order.OrderDataInterface.
func (o *orderQuery) PutUserOrder(userIdLogin int, userOrderId uint, inputOrder order.UserOrder) error {
	putOrder := UserOrderToModel(inputOrder)
	result := o.db.Model(&UserOrder{}).Where("id = ?", userOrderId).Updates(putOrder)
	return result.Error
}

// CheckOrderStatus implements order.OrderDataInterface.
func (o *orderQuery) CheckOrderStatus(userOrderId uint) (string, error) {
	var adminOrder OrderDetail
	result := o.db.Select("status").Where("user_order_id = ?", userOrderId).First(&adminOrder)
	if result.Error != nil {
		return "", result.Error
	}
	return adminOrder.Status, nil
}

// SelectUserOrderWait implements order.OrderDataInterface.
func (o *orderQuery) SelectUserOrderWait(userIdLogin int) ([]order.UserOrder, error) {
	var userOrders []UserOrder

	err := o.db.Preload("User").Preload("OrderDetail").Preload("Region").
		Joins("JOIN order_details ON order_details.user_order_id = user_orders.id").
		Where("user_orders.user_id = ? AND order_details.status = ?", userIdLogin, "Menunggu Diterima").
		Find(&userOrders).Error

	if err != nil {
		return nil, err
	}

	var responseOrders []order.UserOrder
	for _, uo := range userOrders {
		responseOrders = append(responseOrders, uo.ModelToUserOrderWait())
	}

	return responseOrders, nil
}

// SelectById implements order.OrderDataInterface.
func (o *orderQuery) SelectById(IdOrder uint) (*order.UserOrder, error) {
	var userOrderData UserOrder
	err := o.db.Preload("User").
		Preload("Region").
		Preload("OrderDetail").
		First(&userOrderData, IdOrder).Error
	if err != nil {
		log.Printf("Error finding order with ID %d: %v", IdOrder, err)
		return nil, err
	}
	result := userOrderData.ModelToUserOrderWait()
	return &result, nil
}


// InsertOrderDetail implements order.OrderDataInterface.
func (o *orderQuery) InsertOrderDetail(adminIdLogin int, inputOrder order.OrderDetail) error {
	newOrder := OrderDetailToModel(inputOrder)
	adminID := uint(adminIdLogin)
	newOrder.AdminID = &adminID

	result := o.db.Create(&newOrder)
	if result.Error != nil {
		return result.Error
	}

	updateStatus := OrderDetail{
		Status: inputOrder.Status,
	}

	result = o.db.Updates(&updateStatus)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// SelectUserOrderProcess implements order.OrderDataInterface.
func (o *orderQuery) SelectUserOrderProcess(userIdLogin int) ([]order.UserOrder, error) {
	var userOrders []UserOrder

	err := o.db.Preload("User").Preload("OrderDetail").Preload("Region").
		Joins("JOIN order_details ON order_details.user_order_id = user_orders.id").
		Where("user_orders.user_id = ?", userIdLogin).
		Find(&userOrders).Error

	if err != nil {
		return nil, err
	}

	var responseOrders []order.UserOrder
	for _, uo := range userOrders {
		responseOrders = append(responseOrders, uo.ModelToUserOrderWait())
	}

	return responseOrders, nil
}

// SearchUserOrder implements order.OrderDataInterface.
func (o *orderQuery) SearchUserOrder(userIdLogin int, itemName string) ([]order.UserOrder, error) {
	var userOrders []UserOrder

	err := o.db.Preload("User").Preload("OrderDetail").Preload("Region").
		Joins("JOIN order_details ON order_details.user_order_id = user_orders.id").
		Where("user_orders.user_id = ? AND user_orders.item_name LIKE ?", userIdLogin, "%"+itemName+"%").
		Find(&userOrders).Error

	if err != nil {
		return nil, err
	}

	var responseOrders []order.UserOrder
	for _, uo := range userOrders {
		responseOrders = append(responseOrders, uo.ModelToUserOrderWait())
	}

	return responseOrders, nil
}
