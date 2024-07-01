package database

import (
	"fmt"
	"jastip-jakarta/app/config"
	ad "jastip-jakarta/features/admin/data"
	od "jastip-jakarta/features/order/data"
	ud "jastip-jakarta/features/user/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&ud.User{},
		&od.UserOrder{},
		&od.OrderDetail{},
		&ad.Admin{},
		&ad.RegionCode{},
		&ad.DeliveryBatch{},
	)

	return DB
}
