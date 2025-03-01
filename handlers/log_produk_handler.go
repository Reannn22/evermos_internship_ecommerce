package handlers

import (
	"mini-project-evermos/exceptions"
	"mini-project-evermos/middleware"
	"mini-project-evermos/models"
	"mini-project-evermos/models/responder"
	"mini-project-evermos/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductLogHandler struct {
	ProductLogService services.ProductLogService
}

func NewProductLogHandler(productLogService *services.ProductLogService) ProductLogHandler {
	return ProductLogHandler{*productLogService}
}

func (handler *ProductLogHandler) Route(app *fiber.App) {
	routes := app.Group("/api/v1/log-produk")
	routes.Get("/", middleware.JWTProtected(), handler.GetAll)
	routes.Get("/:id", middleware.JWTProtected(), handler.GetById)
	routes.Post("/", middleware.JWTProtected(), handler.Create)
	routes.Put("/:id", middleware.JWTProtected(), handler.Update)
	routes.Delete("/:id", middleware.JWTProtected(), handler.Delete)
}

func (handler *ProductLogHandler) Create(c *fiber.Ctx) error {
	var input models.ProductLogProcess
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to parse request",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.ProductLogService.Create(input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to POST data",
		Error:   nil,
		Data:    response,
	})
}

// Add new handler method
func (handler *ProductLogHandler) GetAll(c *fiber.Ctx) error {
	response, err := handler.ProductLogService.GetAll()
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
		Data:    response,
	})
}

// Add new handler method
func (handler *ProductLogHandler) GetById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.ProductLogService.GetById(uint(id))
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
		Data:    response,
	})
}

// Add new handler method
func (handler *ProductLogHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	var input models.ProductLogProcess
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to parse request",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Attempt update
	response, err := handler.ProductLogService.Update(uint(id), input)
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
				Status:  false,
				Message: "Product log not found",
				Error:   exceptions.NewString(err.Error()),
				Data:    nil,
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to update product log",
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

// Add new handler method
func (handler *ProductLogHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.ProductLogService.Delete(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
				Status:  false,
				Message: "Product log not found",
				Error:   exceptions.NewString(err.Error()),
				Data:    nil,
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to delete product log",
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
