package data

import (
	"jastip-jakarta/features/order"

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

	adminOrder := AdminOrder{
		UserOrderID: newOrder.ID,
		Status:      "Menunggu Diterima",
	}

	result = o.db.Create(&adminOrder)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// PutUserOrder implements order.OrderDataInterface.
func (o *orderQuery) PutUserOrder(userIdLogin int, userOrderId uint, inputOrder order.UserOrder) error {
	result := o.db.Model(&UserOrder{}).Where("id = ?", userOrderId).Updates(inputOrder)
	return result.Error
}

// CheckOrderStatus implements order.OrderDataInterface.
func (o *orderQuery) CheckOrderStatus(userOrderId uint) (string, error) {
	var adminOrder AdminOrder
	result := o.db.Select("status").Where("user_order_id = ?", userOrderId).First(&adminOrder)
	if result.Error != nil {
		return "", result.Error
	}
	return adminOrder.Status, nil
}

// SelectUserOrderWait implements order.OrderDataInterface.
func (o *orderQuery) SelectUserOrderWait(userIdLogin int) ([]order.UserOrder, error) {
	var userOrders []UserOrder

	err := o.db.Preload("User").Preload("AdminOrder").
		Joins("JOIN admin_orders ON admin_orders.user_order_id = user_orders.id").
		Where("user_orders.user_id = ? AND admin_orders.status = ?", userIdLogin, "Menunggu Diterima").
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
