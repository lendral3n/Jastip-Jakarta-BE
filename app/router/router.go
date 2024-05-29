package router

import (
	"jastip-jakarta/utils/cloudinary"
	"jastip-jakarta/utils/encrypts"
	"jastip-jakarta/utils/middlewares"

	ud "jastip-jakarta/features/user/data"
	uh "jastip-jakarta/features/user/handler"
	us "jastip-jakarta/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hash := encrypts.New()
	cloudinaryUploader := cloudinary.New()

	userData := ud.New(db, cloudinaryUploader)
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService)

	// define routes/ endpoint USERS
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/register", userHandlerAPI.RegisterUser)
	e.GET("/users", userHandlerAPI.GetUser, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())
}
