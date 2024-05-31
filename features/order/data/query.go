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
