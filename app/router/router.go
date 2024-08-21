package router

import (
	"jastip-jakarta/utils/cloudinary"
	"jastip-jakarta/utils/csv"
	"jastip-jakarta/utils/encrypts"
	"jastip-jakarta/utils/middlewares"

	ud "jastip-jakarta/features/user/data"
	uh "jastip-jakarta/features/user/handler"
	us "jastip-jakarta/features/user/service"

	ad "jastip-jakarta/features/admin/data"
	ah "jastip-jakarta/features/admin/handler"
	as "jastip-jakarta/features/admin/service"

	od "jastip-jakarta/features/order/data"
	oh "jastip-jakarta/features/order/handler"
	os "jastip-jakarta/features/order/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hash := encrypts.New()
	cloudinaryUploader := cloudinary.New()
	csvGenerator := csv.New()

	userData := ud.New(db, cloudinaryUploader)
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService)

	adminData := ad.New(db, cloudinaryUploader)
	adminService := as.New(adminData, hash, userData)
	adminHandlerAPI := ah.New(adminService)

	orderData := od.New(db, cloudinaryUploader, csvGenerator)
	orderService := os.New(orderData, adminService)
	orderHandlerAPI := oh.New(orderService)

	// define routes/ endpoint USERS
	e.POST("users/login", userHandlerAPI.Login)
	e.POST("users/register", userHandlerAPI.RegisterUser)
	e.GET("/users/profile", userHandlerAPI.GetUser, middlewares.JWTMiddleware())
	e.PUT("/users/profile", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())

	// define routes/ endpoint ADMIN
	e.POST("/admin/register", adminHandlerAPI.RegisterAdminSuper)
	e.POST("/admin/login", adminHandlerAPI.Login)
	e.POST("/admin/new", adminHandlerAPI.RegisterAdmin, middlewares.JWTMiddleware())
	e.GET("/admin/profile", adminHandlerAPI.GetAdmin, middlewares.JWTMiddleware())
	e.PUT("/admin/profile", adminHandlerAPI.UpdateAdmin, middlewares.JWTMiddleware())
	e.GET("/admin/perwakilan", adminHandlerAPI.GetAdminPerwakilan, middlewares.JWTMiddleware())
	e.GET("/admin/jakarta", adminHandlerAPI.GetAdminJakarta, middlewares.JWTMiddleware())
	e.GET("/admin/all", adminHandlerAPI.GetAllAdmin, middlewares.JWTMiddleware())
	e.GET("/admin/user/search", adminHandlerAPI.SearchUser, middlewares.JWTMiddleware())
	e.PUT("/admin/user/:name", adminHandlerAPI.UpdateUserByName, middlewares.JWTMiddleware())
	e.POST("/admin/user", adminHandlerAPI.CreateUser, middlewares.JWTMiddleware())
	e.GET("/admin/user", adminHandlerAPI.GetAllUSer, middlewares.JWTMiddleware())
	
	// define routes/ endpoint BATCH
	e.POST("/admin/batch", adminHandlerAPI.CreateDeliveryBatch, middlewares.JWTMiddleware())
	e.GET("/batch", adminHandlerAPI.GetAllDeliveryBatch)
	e.GET("/batch/:batch_id", adminHandlerAPI.GetDeliveryBatchById)
	
	// define routes/ endpoint REGION
	e.POST("/admin/region", adminHandlerAPI.CreateRegionCode, middlewares.JWTMiddleware())
	e.GET("/region", adminHandlerAPI.GetRegionCode)
	e.GET("/region/:code", adminHandlerAPI.GetRegionCodeById)
	e.GET("/admin/region/search", adminHandlerAPI.SearchRegionCode, middlewares.JWTMiddleware())
	e.PUT("/admin/region/:code", adminHandlerAPI.UpdateRegionCode, middlewares.JWTMiddleware())

	// define routes/ endpoint USER ORDER
	e.POST("/users/order", orderHandlerAPI.CreateUserOrder, middlewares.JWTMiddleware())
	e.PUT("/users/order/:order_id", orderHandlerAPI.UpdateUserOrder, middlewares.JWTMiddleware())
	e.GET("/users/order/wait", orderHandlerAPI.GetUserOrderWait, middlewares.JWTMiddleware())
	e.GET("/users/order/:order_id", orderHandlerAPI.GetOrderById)
	e.GET("/users/order/process", orderHandlerAPI.GetUserOrderProcess, middlewares.JWTMiddleware())
	e.GET("/users/order/search", orderHandlerAPI.SearchUserOrder, middlewares.JWTMiddleware())

	// define routes/ endpoint ADMIN ORDER
	e.POST("/admin/order/:order_id", orderHandlerAPI.CreateOrderDetail, middlewares.JWTMiddleware())
	e.GET("/admin/order", orderHandlerAPI.GetAllUserOrderWait, middlewares.JWTMiddleware())
	e.GET("/admin/order/batch", orderHandlerAPI.GetDeliveryBatchWithRegion, middlewares.JWTMiddleware())
	e.GET("/admin/order/name", orderHandlerAPI.GetUserOrderNames, middlewares.JWTMiddleware())
	e.GET("/admin/order/name/orders", orderHandlerAPI.GetOrderByNameUser, middlewares.JWTMiddleware())
	e.POST("/admin/order/estimasi", orderHandlerAPI.UpdateEstimationForOrders, middlewares.JWTMiddleware())
	e.PUT("/admin/order/status/:order_id", orderHandlerAPI.UpdateOrderStatus, middlewares.JWTMiddleware())
	e.GET("/admin/order/search", orderHandlerAPI.SearchOrder, middlewares.JWTMiddleware())
	e.PUT("/admin/order/:order_id", orderHandlerAPI.UpdateOrderById, middlewares.JWTMiddleware())
	e.GET("/admin/order/statistik/:batch", orderHandlerAPI.GetOrderSStats, middlewares.JWTMiddleware())

	// define routes/ endpoint ADMIN FOTO
	e.POST("/admin/foto", orderHandlerAPI.UploadFotoPacked, middlewares.JWTMiddleware())
	e.PUT("/admin/foto/:id_foto", orderHandlerAPI.UploadFotoReceived, middlewares.JWTMiddleware())

	// define routes/ endpoint ADMIN CSV
	e.GET("/download/csv", orderHandlerAPI.GenerateCSVByBatch)
}