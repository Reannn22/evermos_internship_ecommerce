package handlers

import (
	"fmt"
	"mini-project-evermos/exceptions"
	"mini-project-evermos/middleware"
	"mini-project-evermos/models"
	"mini-project-evermos/models/responder"
	"mini-project-evermos/services"
	"mini-project-evermos/utils/jwt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductPromoHandler struct {
	service services.ProductPromoService
}

func NewProductPromoHandler(service *services.ProductPromoService) ProductPromoHandler {
	return ProductPromoHandler{*service}
}

func (h *ProductPromoHandler) Route(app *fiber.App) {
	routes := app.Group("/api/v1/promos")
	routes.Get("/", h.GetAll)
	routes.Delete("/clear", middleware.JWTProtected(), h.ClearAll) // Move this route before the /:id routes
	routes.Get("/:id", h.GetById)
	routes.Get("/product/:productId", h.GetByProductId)
	routes.Post("/", middleware.JWTProtected(), h.Create)
	routes.Put("/:id", middleware.JWTProtected(), h.Update)
	routes.Delete("/:id", middleware.JWTProtected(), h.Delete)
}

func (h *ProductPromoHandler) GetAll(c *fiber.Ctx) error {
	promos, err := h.service.GetAll()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get promos",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully retrieved promos",
		Error:   nil,
		Data:    promos,
	})
}

func (h *ProductPromoHandler) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	promo, err := h.service.GetById(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get promo",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully retrieved promo",
		Error:   nil,
		Data:    promo,
	})
}

func (h *ProductPromoHandler) GetByProductId(c *fiber.Ctx) error {
	productId, err := strconv.ParseUint(c.Params("productId"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid product ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	promos, err := h.service.GetByProductId(uint32(productId))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get promos",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully retrieved promos",
		Error:   nil,
		Data:    promos,
	})
}

func (h *ProductPromoHandler) Create(c *fiber.Ctx) error {
	// Add debug logging
	fmt.Println("Received request to create promo")

	var input models.ProductPromoRequest
	if err := c.BodyParser(&input); err != nil {
		fmt.Printf("Error parsing request body: %v\n", err)
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid request body",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	fmt.Printf("Parsed request body: %+v\n", input)

	promo, err := h.service.Create(input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to create promo",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully created promo",
		Error:   nil,
		Data:    promo,
	})
}

func (h *ProductPromoHandler) Update(c *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	var input models.ProductPromoRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid request body",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	promo, err := h.service.Update(input, id, uint64(claims.UserId))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to update promo",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully updated promo",
		Error:   nil,
		Data:    promo,
	})
}

func (h *ProductPromoHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Get the promo data before deletion
	promo, err := h.service.GetById(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get promo",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	err = h.service.Delete(id, 0)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to delete promo",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully deleted promo",
		Error:   nil,
		Data:    promo,
	})
}

// Add this new method at the end of the file
func (h *ProductPromoHandler) ClearAll(c *fiber.Ctx) error {
	promos, err := h.service.ClearAll()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to clear all promos",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully cleared all promos",
		Error:   nil,
		Data:    promos,
	})
}
