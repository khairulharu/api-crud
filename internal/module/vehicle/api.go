package vehicle

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/miniapps/domain"
	"github.com/khairulharu/miniapps/internal/util"
)

type api struct {
	vehicleService domain.VehicleService
}

func NewApi(app *fiber.App, vehicleService domain.VehicleService) {
	api := api{
		vehicleService: vehicleService,
	}

	app.Get("v1/vehicle-histories", api.GetVehicleHistories)
	app.Post("v1/vehicle-histories", api.StoreVehicleHistori)
}

func (a api) GetVehicleHistories(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	vin := ctx.Query("vin")
	if vin == "" {
		apiResponse := domain.ApiResponse{
			Code:    "01",
			Message: "INVALID PARAMETER",
		}
		util.ResponseInterceptor(c, &apiResponse)

		return ctx.Status(400).JSON(apiResponse)
	}

	apiResponse := a.vehicleService.FindHistorical(c, vin)
	util.ResponseInterceptor(c, &apiResponse)

	return ctx.Status(200).JSON(apiResponse)
}

func (a api) StoreVehicleHistori(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	var req domain.VehicleHistoricalRequest

	if err := ctx.BodyParser(&req); err != nil {
		apiResponse := domain.ApiResponse{
			Code:    "911",
			Message: "INVALID BODY JSON PARSER",
		}
		util.ResponseInterceptor(c, &apiResponse)

		return ctx.Status(400).JSON(apiResponse)
	}

	apiResponse := a.vehicleService.StoreHistorical(c, req)
	util.ResponseInterceptor(c, &apiResponse)

	return ctx.Status(200).JSON(apiResponse)
}
