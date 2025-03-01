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

type NotificationHandler struct {
	service services.NotificationService
}

func NewNotificationHandler(service services.NotificationService) *NotificationHandler {
	return &NotificationHandler{service}
}

func (h *NotificationHandler) Route(app *fiber.App) {
	routes := app.Group("/api/v1/notifications")
	routes.Use(middleware.JWTProtected())
	routes.Get("/", h.GetAll)
	routes.Get("/:id", h.GetById)
	routes.Post("/", h.Create)
	routes.Put("/:id", h.Update)
	routes.Delete("/:id", h.Delete)
}

func (h *NotificationHandler) GetAll(c *fiber.Ctx) error {
	notifications, err := h.service.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get notifications",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success get notifications",
		Data:    notifications,
	})
}

func (h *NotificationHandler) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	notification, err := h.service.GetById(uint(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Notification not found",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success get notification",
		Data:    notification,
	})
}

func (h *NotificationHandler) Create(c *fiber.Ctx) error {
	var input models.NotificationRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid request body",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	notification, err := h.service.Create(input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to create notification",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	return c.Status(http.StatusCreated).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success create notification",
		Data:    notification,
	})
}

func (h *NotificationHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	var input models.NotificationRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid request body",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	notification, err := h.service.Update(uint(id), input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to update notification",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success update notification",
		Data:    notification,
	})
}

func (h *NotificationHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	notification, err := h.service.Delete(uint(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to delete notification",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success delete notification",
		Data:    notification,
	})
}
