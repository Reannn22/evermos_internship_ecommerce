package handlers

import (
	"mini-project-evermos/exceptions"
	"mini-project-evermos/middleware"
	"mini-project-evermos/models"
	"mini-project-evermos/models/responder"
	"mini-project-evermos/services"
	"mini-project-evermos/utils/jwt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FotoProdukHandler struct {
	service services.FotoProdukService
}

func NewFotoProdukHandler(service *services.FotoProdukService) FotoProdukHandler {
	return FotoProdukHandler{*service}
}

func (handler *FotoProdukHandler) Route(app *fiber.App) {
	routes := app.Group("/api/v1/foto-produk")
	routes.Get("/", middleware.JWTProtected(), handler.GetAll)
	routes.Get("/:id", middleware.JWTProtected(), handler.GetById)
	routes.Get("/product/:id", middleware.JWTProtected(), handler.GetByProductId)
	routes.Post("/", middleware.JWTProtected(), handler.Create)
	routes.Put("/:id", middleware.JWTProtected(), handler.Update)
	routes.Delete("/:id", middleware.JWTProtected(), handler.Delete)
}

func (handler *FotoProdukHandler) GetAll(c *fiber.Ctx) error {
	photos, err := handler.service.FindAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get photos",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success get photos",
		Error:   nil,
		Data:    photos,
	})
}

func (handler *FotoProdukHandler) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	photo, err := handler.service.FindById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get photo",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success get photo",
		Error:   nil,
		Data:    photo,
	})
}

func (handler *FotoProdukHandler) GetByProductId(c *fiber.Ctx) error {
	productId, err := strconv.ParseUint(c.Params("product_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid product ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	photos, err := handler.service.FindByProductId(uint(productId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get photos",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success get photos",
		Error:   nil,
		Data:    photos,
	})
}

func (handler *FotoProdukHandler) Create(c *fiber.Ctx) error {
	_, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Parse form
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get file",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Convert id_produk from string to uint
	productID, err := strconv.ParseUint(c.FormValue("id_produk", "0"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid product ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Create request object
	request := models.FotoProdukRequest{
		ProductID: uint(productID),
		PhotoURL:  c.FormValue("photo_url", ""),
		File:      file,
	}

	// Save the file
	filename := filepath.Base(file.Filename)
	if err := c.SaveFile(file, "./uploads/products/"+filename); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to save file",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Call service
	response, err := handler.service.Create(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to create photo",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success create photo",
		Error:   nil,
		Data:    response,
	})
}

func (handler *FotoProdukHandler) Update(c *fiber.Ctx) error {
	_, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Parse form
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get file",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Convert id_produk from string to uint
	productID, err := strconv.ParseUint(c.FormValue("id_produk", "0"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid product ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Create request object
	request := models.FotoProdukRequest{
		ProductID: uint(productID),
		PhotoURL:  c.FormValue("photo_url", ""),
		File:      file,
	}

	// Save the file
	filename := filepath.Base(file.Filename)
	if err := c.SaveFile(file, "./uploads/products/"+filename); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to save file",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.service.Update(uint(id), request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to update photo",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success update photo",
		Error:   nil,
		Data:    response,
	})
}

func (handler *FotoProdukHandler) Delete(c *fiber.Ctx) error {
	_, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	deletedPhoto, err := handler.service.Delete(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to delete photo",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success delete photo",
		Error:   nil,
		Data:    deletedPhoto,
	})
}
