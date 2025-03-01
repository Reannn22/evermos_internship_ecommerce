package handlers

import (
	"mini-project-evermos/exceptions"
	"mini-project-evermos/middleware"
	"mini-project-evermos/models"
	"mini-project-evermos/models/responder"
	"mini-project-evermos/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type DiskonProdukHandler struct {
	DiskonProdukService services.DiskonProdukService
}

func NewDiskonProdukHandler(diskonProdukService services.DiskonProdukService) *DiskonProdukHandler {
	return &DiskonProdukHandler{DiskonProdukService: diskonProdukService}
}

func (handler *DiskonProdukHandler) Route(app *fiber.App) {
	route := app.Group("/api/v1/diskon-produk")
	route.Post("/", middleware.JWTProtected(), handler.ApplyDiscount)
	route.Get("/", middleware.JWTProtected(), handler.GetAllDiscounts)
	route.Get("/:id", middleware.JWTProtected(), handler.GetDiscountById)
	route.Put("/:id", middleware.JWTProtected(), handler.UpdateDiscount)
	route.Delete("/:id", middleware.JWTProtected(), handler.DeleteDiscount) // Add this line
}

func (handler *DiskonProdukHandler) ApplyDiscount(c *fiber.Ctx) error {
	var input models.DiskonProdukRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to parse request",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.DiskonProdukService.ApplyDiscount(input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to apply discount",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Use struct tags to maintain field order
	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to POST data",
		Error:   nil,
		Data:    response, // The struct tags will handle the field order
	})
}

func (handler *DiskonProdukHandler) GetAllDiscounts(c *fiber.Ctx) error {
	responses, err := handler.DiskonProdukService.GetAll()
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

func (handler *DiskonProdukHandler) GetDiscountById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.DiskonProdukService.GetById(uint(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Discount not found",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Error:   nil,
		Data:    response,
	})
}

func (handler *DiskonProdukHandler) UpdateDiscount(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	var input models.DiskonProdukRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to parse request",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.DiskonProdukService.UpdateDiscount(uint(id), input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to update discount",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to PUT data",
		Error:   nil,
		Data:    response,
	})
}

// Add this new method
func (handler *DiskonProdukHandler) DeleteDiscount(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.DiskonProdukService.DeleteDiscount(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to delete discount",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to DELETE data",
		Error:   nil,
		Data:    response,
	})
}
