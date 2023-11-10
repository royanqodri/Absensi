package router

import (
	"Absensi-App/app/middlewares"

	dataA "Absensi-App/features/absensi/data"
	handlerA "Absensi-App/features/absensi/handler"
	serviceA "Absensi-App/features/absensi/service"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(c *echo.Echo, db *gorm.DB, redis *redis.Client) {
	externalAPI := apinodejs.NewExternalData("https://project2.otixx.online")

	dataAbsensi := dataA.New(db, externalAPI)
	serviceAbsensi := serviceA.New(dataAbsensi)
	handlerAbsensi := handlerA.New(serviceAbsensi)

	c.POST("/absensis", handlerAbsensi.Add, middlewares.JWTMiddleware())
	c.PUT("/absensis/:id_absensi", handlerAbsensi.Edit, middlewares.JWTMiddleware())
	c.GET("/absensis", handlerAbsensi.GetAll, middlewares.JWTMiddleware())
	c.GET("/absensis/:id_absensi", handlerAbsensi.GetAbsensiById, middlewares.JWTMiddleware())
}
