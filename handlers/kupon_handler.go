package handlers

import (
	"mini-project-evermos/exceptions"
	"mini-project-evermos/middleware"
	"mini-project-evermos/models"
	"mini-project-evermos/models/responder"
	"mini-project-evermos/services"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductCouponHandler struct {
	service services.ProductCouponService
}

func NewProductCouponHandler(service services.ProductCouponService) *ProductCouponHandler {
	return &ProductCouponHandler{service}
}

func (h *ProductCouponHandler) Route(app *fiber.App) {
	routes := app.Group("/api/v1/coupons")
	routes.Use(middleware.JWTProtected())

	routes.Get("/", h.GetAll)
	routes.Get("/:id", h.GetById)
	routes.Get("/code/:code", h.GetByCode)
	routes.Get("/product/:productId", h.GetByProduct)
	routes.Post("/", h.Create)
	routes.Put("/:id", h.Update)
	routes.Delete("/:id", h.Delete)
	routes.Post("/validate", h.ValidateCoupon)
}

func (h *ProductCouponHandler) GetAll(c *fiber.Ctx) error {
	coupons, err := h.service.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get coupons",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully retrieved coupons",
		Error:   nil,
		Data:    coupons,
	})
}

func (h *ProductCouponHandler) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	coupon, err := h.service.GetById(uint(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Coupon not found",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully retrieved coupon",
		Error:   nil,
		Data:    coupon,
	})
}

func (h *ProductCouponHandler) GetByCode(c *fiber.Ctx) error {
	code := c.Params("code")
	coupon, err := h.service.GetByCode(code)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Coupon not found",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully retrieved coupon",
		Error:   nil,
		Data:    coupon,
	})
}

func (h *ProductCouponHandler) GetByProduct(c *fiber.Ctx) error {
	productId, err := strconv.ParseUint(c.Params("productId"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid product ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	coupons, err := h.service.GetByProduct(uint(productId))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get coupons",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully retrieved coupons",
		Error:   nil,
		Data:    coupons,
	})
}

func (h *ProductCouponHandler) Create(c *fiber.Ctx) error {
	var request models.ProductCouponRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid request body",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	coupon, err := h.service.Create(request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to create coupon",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully created coupon",
		Error:   nil,
		Data:    coupon,
	})
}

func (h *ProductCouponHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	var request models.ProductCouponRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid request body",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	coupon, err := h.service.Update(uint(id), request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to update coupon",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully updated coupon",
		Error:   nil,
		Data:    coupon,
	})
}

func (h *ProductCouponHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	deletedCoupon, err := h.service.Delete(uint(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to delete coupon",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully deleted coupon",
		Error:   nil,
		Data:    deletedCoupon,
	})
}

func (h *ProductCouponHandler) ValidateCoupon(c *fiber.Ctx) error {
	type ValidateRequest struct {
		Code      string `json:"code"`
		ProductID uint   `json:"product_id"`
	}

	var request ValidateRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid request body",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	coupon, err := h.service.ValidateCoupon(request.Code, request.ProductID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid or expired coupon",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Coupon is valid",
		Error:   nil,
		Data:    coupon,
	})
}
