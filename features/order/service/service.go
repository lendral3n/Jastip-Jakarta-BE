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
func (o *orderService) CreateUserOrder(userIdLogin int, inputOrder order.UserOrder) error {
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

// SelectUserOrderWait implements order.OrderServiceInterface.
func (o *orderService) GetUserOrderWait(userIdLogin int) ([]order.UserOrder, error) {
	userOrders, err := o.orderData.SelectUserOrderWait(userIdLogin)
	if err != nil {
		return nil, err
	}
	return userOrders, nil
}

// GetById implements order.OrderServiceInterface.
func (o *orderService) GetById(IdOrder uint) (*order.UserOrder, error) {
	result, err := o.orderData.SelectById(IdOrder)
	return result, err
}

// CreateAdminOrder implements order.OrderServiceInterface.
func (o *orderService) CreateAdminOrder(adminIdLogin int, userOrderId uint, inputOrder order.AdminOrder) error {
	orderIdCheck, err := o.orderData.SelectById(userOrderId)
	if err != nil {
		return err
	}

	if orderIdCheck.ID != userOrderId {
		return errors.New("ID Order tidak ditemukan atau salah")
	}

	if inputOrder.Status == "" {
		return errors.New("Status Harus Di Isi")
	}

	if inputOrder.WeightItem == 0 {
		return errors.New("Berat Tidak Boleh Nol")
	}

	if inputOrder.DeliveryBatch == "" {
		return errors.New("Batch Pengiriman Tidak Boleh Kosong")
	}

	if inputOrder.TrackingNumberJastip == "" {
		return errors.New("No Resi JASTIP Tidak Boleh Kosong")
	}

	return o.orderData.InsertAdminOrder(adminIdLogin, inputOrder)
}

// GetUserOrderProcess implements order.OrderServiceInterface.
func (o *orderService) GetUserOrderProcess(userIdLogin int) ([]order.UserOrder, error) {
	userOrders, err := o.orderData.SelectUserOrderProcess(userIdLogin)
	if err != nil {
		return nil, err
	}
	return userOrders, nil
}