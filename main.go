package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/khairulharu/miniapps/internal/component"
	"github.com/khairulharu/miniapps/internal/config"
	"github.com/khairulharu/miniapps/internal/module/customer"
	"github.com/khairulharu/miniapps/internal/module/history"
	"github.com/khairulharu/miniapps/internal/module/vehicle"
)

func main() {
	conf := config.Get()
	dbConnection := component.GetDatabase(conf)

	customerRepository := customer.NewRepository(dbConnection)
	vehicleRepository := vehicle.NewRepository(dbConnection)
	historyRepository := history.NewRepository(dbConnection)

	customerService := customer.NewService(customerRepository)
	vehicleService := vehicle.NewService(historyRepository, vehicleRepository)

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${locals:requestid}] ${ip} - ${method} ${status} ${path}\n",
	}))
	customer.NewApi(app, customerService)
	vehicle.NewApi(app, vehicleService)
	_ = app.Listen(conf.SRV.Host + ":" + conf.SRV.Port)
}
