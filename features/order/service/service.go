package service

import (
	"errors"
	"jastip-jakarta/features/order"
)

type orderService struct {
	orderData order.OrderDataInterface
}

func New(repo order.OrderDataInterface) order.OrderServiceInterface {
	return &orderService{
		orderData: repo,
	}
}

// CreateOrder implements order.OrderServiceInterface.
func (o *orderService) CreateOrder(userIdLogin int, inputOrder order.UserOrder) error {
	if inputOrder.ItemName == "" {
		return errors.New("Nama Barang harus diisi")
	}
	if inputOrder.TrackingNumber == "" {
		return errors.New("Nomor resi harus diisi")
	}
	if inputOrder.OnlineStore == "" {
		return errors.New("Nama toko online harus diisi")
	}
	if inputOrder.WhatsAppNumber == 0 {
		return errors.New("Nomor WhatsApp harus diisi")
	}
	if inputOrder.RegionCode == "" {
		return errors.New("Kode wilayah harus diisi")
	}

	return o.orderData.InsertUserOrder(userIdLogin, inputOrder)
}

// UpdateUserOrder implements order.OrderServiceInterface.
func (o *orderService) UpdateUserOrder(userIdLogin int, userOrderId uint, inputOrder order.UserOrder) error {
	// Mengecek status terlebih dahulu
	status, err := o.orderData.CheckOrderStatus(userOrderId)
	if err != nil {
		return err
	}

	// Melakukan update jika status 'Menunggu Diterima'
	if status == "Menunggu Diterima" {
		err = o.orderData.PutUserOrder(userIdLogin, userOrderId, inputOrder)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Order tidak dapat diupdate karena status bukan 'Menunggu Diterima'")
	}

	return nil
}