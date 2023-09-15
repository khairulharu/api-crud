package customer

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/miniapps/domain"
	"github.com/khairulharu/miniapps/internal/util"
)

type api struct {
	customerService domain.CustomerService
}

func NewApi(app *fiber.App, customerService domain.CustomerService) {
	api := api{
		customerService,
	}

	app.Get("/v1/customers", api.AllCustomers)
	app.Post("/v1/customers", api.SaveCustomer)

}

func (a api) AllCustomers(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	apiResponse := a.customerService.All(c)
	util.ResponseInterceptor(c, &apiResponse)

	return ctx.Status(200).JSON(apiResponse)
}

func (a api) SaveCustomer(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	var customerData domain.CustomerData

	if err := ctx.BodyParser(&customerData); err != nil {
		apiResponse := domain.ApiResponse{
			Code:    "911",
			Message: "BODY PARSER FAILED TO REQUSET JSON",
		}
		util.ResponseInterceptor(c, &apiResponse)
		ctx.Status(400).JSON(apiResponse)
	}

	apiResponse := a.customerService.Save(c, customerData)
	util.ResponseInterceptor(c, &apiResponse)

	return ctx.Status(200).JSON(apiResponse)
}
