package handlers

import (
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

type ProductReviewHandler struct {
	service services.ProductReviewService
}

func NewProductReviewHandler(service *services.ProductReviewService) ProductReviewHandler {
	return ProductReviewHandler{*service}
}

func (h *ProductReviewHandler) Route(app *fiber.App) {
	routes := app.Group("/api/v1/reviews")
	routes.Get("/", h.GetAll)
	routes.Get("/:id", h.GetById)
	routes.Get("/product/:productId", h.GetByProductId)
	routes.Post("/", middleware.JWTProtected(), h.Create)
	routes.Put("/:id", middleware.JWTProtected(), h.Update)
	routes.Delete("/:id", middleware.JWTProtected(), h.Delete)
}

func (h *ProductReviewHandler) GetAll(c *fiber.Ctx) error {
	reviews, err := h.service.GetAll()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get reviews",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully retrieved reviews",
		Error:   nil,
		Data:    reviews,
	})
}

func (h *ProductReviewHandler) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	review, err := h.service.GetById(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get review",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully retrieved review",
		Error:   nil,
		Data:    review,
	})
}

func (h *ProductReviewHandler) GetByProductId(c *fiber.Ctx) error {
	productId, err := strconv.ParseUint(c.Params("productId"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid product ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	reviews, err := h.service.GetByProductId(uint32(productId))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get reviews",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully retrieved reviews",
		Error:   nil,
		Data:    reviews,
	})
}

func (h *ProductReviewHandler) Create(c *fiber.Ctx) error {
	var input models.ProductReviewRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid request body",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	review, err := h.service.Create(input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to create review",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully created review",
		Error:   nil,
		Data:    review,
	})
}

func (h *ProductReviewHandler) Update(c *fiber.Ctx) error {
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

	var input models.ProductReviewRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid request body",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	review, err := h.service.Update(input, id, uint64(claims.UserId))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to update review",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully updated review",
		Error:   nil,
		Data:    review,
	})
}

func (h *ProductReviewHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Get the review data before deletion
	review, err := h.service.GetById(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get review",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Delete the review
	err = h.service.Delete(id, 0) // Pass 0 as userId since we're not checking it anymore
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to delete review",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully deleted review",
		Error:   nil,
		Data:    review,
	})
}
