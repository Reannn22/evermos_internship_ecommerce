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

type KeranjangBelanjaHandler struct {
	KeranjangBelanjaService services.KeranjangBelanjaService
}

func NewKeranjangBelanjaHandler(keranjangBelanjaService *services.KeranjangBelanjaService) KeranjangBelanjaHandler {
	return KeranjangBelanjaHandler{*keranjangBelanjaService}
}

func (handler *KeranjangBelanjaHandler) Route(app *fiber.App) {
	routes := app.Group("/api/v1/keranjang-belanja")
	routes.Use(middleware.JWTProtected())

	// Place the /clear route before the /:id route to prevent parameter confusion
	routes.Delete("/clear", handler.ClearAll)

	routes.Get("/", handler.GetAll)
	routes.Get("/:id", handler.GetById)
	routes.Post("/", handler.Create)
	routes.Put("/:id", handler.Update)
	routes.Delete("/:id", handler.Delete)
}

func (handler *KeranjangBelanjaHandler) GetAll(c *fiber.Ctx) error {
	responses, err := handler.KeranjangBelanjaService.GetAll()
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
		Message: "Success to GET data",
		Error:   nil,
		Data:    responses,
	})
}

func (handler *KeranjangBelanjaHandler) GetById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.KeranjangBelanjaService.GetById(uint(id))
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
		Message: "Success to GET data",
		Error:   nil,
		Data:    response,
	})
}

func (handler *KeranjangBelanjaHandler) Create(c *fiber.Ctx) error {
	input := new(models.KeranjangBelanjaRequest)
	if err := c.BodyParser(input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to parse request",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.KeranjangBelanjaService.Create(*input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to CREATE data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success to CREATE data",
		Error:   nil,
		Data:    response,
	})
}

func (handler *KeranjangBelanjaHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	input := new(models.KeranjangBelanjaRequest)
	if err := c.BodyParser(input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to parse request",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.KeranjangBelanjaService.Update(uint(id), *input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to UPDATE data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success to UPDATE data",
		Error:   nil,
		Data:    response,
	})
}

func (handler *KeranjangBelanjaHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.KeranjangBelanjaService.Delete(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to DELETE data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success to DELETE data",
		Error:   nil,
		Data:    response,
	})
}

func (handler *KeranjangBelanjaHandler) ClearAll(c *fiber.Ctx) error {
	deletedItems, err := handler.KeranjangBelanjaService.ClearAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to clear shopping cart",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully cleared shopping cart",
		Error:   nil,
		Data:    deletedItems, // Return the deleted items
	})
}
