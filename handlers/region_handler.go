package handlers

import (
	"mini-project-evermos/exceptions"
	"mini-project-evermos/models/responder"
	"mini-project-evermos/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type RegionHandler struct {
	RegionService services.RegionService
}

func NewRegionHandler(regionService *services.RegionService) RegionHandler {
	return RegionHandler{*regionService}
}

// Route
func (handler *RegionHandler) Route(app *fiber.App) {
	routes := app.Group("/api/v1/provcity")
	routes.Get("/listprovincies", handler.Provicies)
	routes.Get("/listcities/:prov_id", handler.Cities)
	routes.Get("/detailprovince/:prov_id", handler.Province)
	routes.Get("/detailcity/:city_id", handler.City)
}

func (handler *RegionHandler) Provicies(c *fiber.Ctx) error {
	responses, err := handler.RegionService.GetAllProvince()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Error:   nil,
		Data:    responses,
	})
}

func (handler *RegionHandler) Province(c *fiber.Ctx) error {
	prov_id := c.Params("prov_id")

	responses, err := handler.RegionService.GetProvince(prov_id)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Error:   nil,
		Data:    responses,
	})
}

func (handler *RegionHandler) Cities(c *fiber.Ctx) error {
	prov_id := c.Params("prov_id")

	responses, err := handler.RegionService.GetAllCity(prov_id)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Error:   nil,
		Data:    responses,
	})
}

func (handler *RegionHandler) City(c *fiber.Ctx) error {
	city_id := c.Params("city_id")

	responses, err := handler.RegionService.GetCity(city_id)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Error:   nil,
		Data:    responses,
	})
}
