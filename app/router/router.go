package router

import (
	"jastip-jakarta/utils/cloudinary"
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

	userData := ud.New(db, cloudinaryUploader)
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService)

	adminData := ad.New(db, cloudinaryUploader)
	adminService := as.New(adminData, hash)
	adminHandlerAPI := ah.New(adminService)

	orderData := od.New(db)
	orderService := os.New(orderData)
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

	// define routes/ endpoint USER ORDER
	e.POST("/users/order", orderHandlerAPI.CreateUserOrder, middlewares.JWTMiddleware())
	e.PUT("/users/order/:order_id", orderHandlerAPI.UpdateUserOrder, middlewares.JWTMiddleware())
	e.GET("/users/order/wait", orderHandlerAPI.GetUserOrderWait, middlewares.JWTMiddleware())
}
