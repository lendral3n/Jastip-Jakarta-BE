package router

import (
	"jastip-jakarta/utils/cloudinary"
	"jastip-jakarta/utils/encrypts"
	"jastip-jakarta/utils/middlewares"

	ud "jastip-jakarta/features/user/data"
	uh "jastip-jakarta/features/user/handler"
	us "jastip-jakarta/features/user/service"

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

	orderData := od.New(db)
	orderService := os.New(orderData)
	orderHandlerAPI := oh.New(orderService)

	// define routes/ endpoint USERS
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/register", userHandlerAPI.RegisterUser)
	e.GET("/users", userHandlerAPI.GetUser, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())

	// define routes/ endpoint USER ORDER
	e.POST("/users/order", orderHandlerAPI.CreateUserOrder, middlewares.JWTMiddleware())
	e.PUT("/users/order/:order_id", orderHandlerAPI.UpdateUserOrder, middlewares.JWTMiddleware())
	e.GET("/users/order/wait", orderHandlerAPI.GetUserOrderWait, middlewares.JWTMiddleware())
}
