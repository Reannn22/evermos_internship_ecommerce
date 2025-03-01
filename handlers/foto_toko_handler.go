package handlers

import (
	"fmt"
	"mini-project-evermos/exceptions"
	"mini-project-evermos/middleware"
	"mini-project-evermos/models"
	"mini-project-evermos/models/responder"
	"mini-project-evermos/services"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type StorePhotoHandler struct {
	StorePhotoService services.StorePhotoService
}

func NewStorePhotoHandler(service *services.StorePhotoService) StorePhotoHandler {
	return StorePhotoHandler{*service}
}

func (handler *StorePhotoHandler) Route(app *fiber.App) {
	routes := app.Group("/api/v1/toko-foto")
	routes.Use(middleware.JWTProtected())
	routes.Get("/", handler.GetAllPhotos) // Add this new route
	routes.Get("/:id_toko", handler.GetStorePhotos)
	routes.Post("/", handler.CreatePhoto)
	routes.Put("/:id", handler.UpdatePhoto)
	routes.Delete("/:id", handler.DeletePhoto)
}

// Add this new handler function
func (handler *StorePhotoHandler) GetAllPhotos(c *fiber.Ctx) error {
	responses, err := handler.StorePhotoService.GetAll()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Return all photos instead of just the first one
	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Error:   nil,
		Data:    responses, // Changed from responses[0] to responses to return all photos
	})
}

func (handler *StorePhotoHandler) GetStorePhotos(c *fiber.Ctx) error {
	// Get the photo ID from the URL parameter
	photoId, err := c.ParamsInt("id_toko")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid photo ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Get single photo by ID
	photo, err := handler.StorePhotoService.GetById(uint(photoId))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get photo",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Error:   nil,
		Data:    photo,
	})
}

func (handler *StorePhotoHandler) CreatePhoto(c *fiber.Ctx) error {
	idTokoStr := strings.TrimSpace(c.FormValue("id_toko"))
	idToko, err := strconv.Atoi(idTokoStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid id_toko",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Change "photo" to "foto" to match the form-data key
	file, err := c.FormFile("foto")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Photo file is required",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Generate filename with timestamp
	timestamp := time.Now().Format("20060102150405")
	filename := fmt.Sprintf("%s_%s", timestamp, file.Filename)
	photo_path := fmt.Sprintf("/uploads/%s", filename)

	// Save the file
	err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", filename))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to save uploaded file",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	input := models.StorePhotoRequest{
		IdToko: uint(idToko),
		URL:    c.FormValue("url"),
		Photo:  photo_path,
		IdFoto: strconv.Itoa(idToko), // Add this line to set IdFoto
	}

	response, err := handler.StorePhotoService.Create(input)
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

func (handler *StorePhotoHandler) UpdatePhoto(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Get form values correctly
	idToko, err := strconv.Atoi(c.FormValue("id_toko", "0")) // Fix this line
	if err != nil {
		idToko = 0 // Default value if conversion fails
	}

	input := models.StorePhotoRequest{
		IdToko: uint(idToko),
		IdFoto: c.FormValue("id_foto"),
		URL:    c.FormValue("url"),
	}

	// Handle file upload if provided
	if file, err := c.FormFile("photo"); err == nil {
		timestamp := time.Now().Format("20060102150405")
		filename := fmt.Sprintf("%s_%s", timestamp, file.Filename)

		err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", filename))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
				Status:  false,
				Message: "Failed to save uploaded file",
				Error:   exceptions.NewString(err.Error()),
				Data:    nil,
			})
		}

		input.Photo = fmt.Sprintf("/uploads/%s", filename)
	}

	response, err := handler.StorePhotoService.Update(uint(id), input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to update photo",
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

func (handler *StorePhotoHandler) DeletePhoto(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to DELETE data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Get the photo data before deletion
	deletedPhoto, err := handler.StorePhotoService.Delete(uint(id))
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
		Message: "Succeed to DELETE data",
		Error:   nil,
		Data:    deletedPhoto,
	})
}
