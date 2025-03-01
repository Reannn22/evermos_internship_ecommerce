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
	"time"

	"github.com/gofiber/fiber/v2"
)

type StoreHandler struct {
	StoreService services.StoreService
}

func NewStoreHandler(storeService *services.StoreService) StoreHandler {
	return StoreHandler{*storeService}
}

func (handler *StoreHandler) Route(app *fiber.App) {
	fmt.Println("Registering store routes") // Add debug log
	routes := app.Group("/api/v1/toko")

	// GET endpoints - accessible by all authenticated users
	routes.Get("/", middleware.JWTProtected(), handler.GetAllStore)
	routes.Get("/my", middleware.JWTProtected(), handler.MyStore)
	routes.Get("/:id_toko", middleware.JWTProtected(), handler.StoreDetail)

	// POST/PUT/DELETE endpoints - admin only
	routes.Post("/", middleware.JWTProtected(), handler.adminOnly, handler.StoreCreate)
	routes.Put("/:id_toko", middleware.JWTProtected(), handler.adminOnly, handler.EditStore)
	routes.Delete("/:id_toko", middleware.JWTProtected(), handler.adminOnly, handler.DeleteStore)
}

// Add new middleware function to check admin status
func (handler *StoreHandler) adminOnly(c *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	if !claims.IsAdmin {
		return c.Status(http.StatusForbidden).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Access denied: Admin only",
			Error:   exceptions.NewString("forbidden access"),
			Data:    nil,
		})
	}

	return c.Next()
}

func (handler *StoreHandler) MyStore(c *fiber.Ctx) error {
	//claim
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	user_id := claims.UserId

	responses, err := handler.StoreService.GetByUserId(uint(user_id))
	if err != nil {
		//error
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

func (handler *StoreHandler) GetAllStore(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(c.FormValue("limit"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   exceptions.NewString("limit required."),
			Data:    nil,
		})
	}

	page, err := strconv.Atoi(c.FormValue("page"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   exceptions.NewString("page required."),
			Data:    nil,
		})
	}

	keyword := c.FormValue("nama")

	responses, err := handler.StoreService.GetAll(limit, page, keyword)

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

func (handler *StoreHandler) StoreDetail(c *fiber.Ctx) error {
	//claim
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	user_id := claims.UserId

	id, err := c.ParamsInt("id_toko")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.StoreService.GetById(uint(id), uint(user_id))
	if err != nil {
		//error
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

func (handler *StoreHandler) EditStore(c *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to PUT data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	user_id := claims.UserId
	id_toko, err := c.ParamsInt("id_toko")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to PUT data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	name_store := c.FormValue("nama_toko")
	deskripsi_toko := c.FormValue("deskripsi_toko")

	// Handle file upload
	var photo_path string
	file, err := c.FormFile("photo")
	if err == nil {
		// If file was uploaded successfully
		timestamp := time.Now().Format("20060102150405")
		filename := fmt.Sprintf("%s_%s", timestamp, file.Filename)
		photo_path = fmt.Sprintf("/uploads/%s", filename)

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
	}

	url_foto := c.FormValue("url_foto")
	if url_foto == "" && file != nil {
		url_foto = file.Filename
	} else if url_foto == "" {
		url_foto = "default-store-image.jpg"
	}

	input := models.StoreProcess{
		ID:            uint(id_toko),
		UserID:        uint(user_id),
		NamaToko:      name_store,
		DeskripsiToko: deskripsi_toko,
		URL:           url_foto,
		Photo:         photo_path,
	}

	response, err := handler.StoreService.Edit(input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to PUT data",
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

func (handler *StoreHandler) StoreCreate(c *fiber.Ctx) error {
	fmt.Println("StoreCreate handler called")
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	user_id := claims.UserId
	id_user_str := c.FormValue("id_user") // Add this line
	nama_toko := c.FormValue("nama_toko")
	deskripsi_toko := c.FormValue("deskripsi_toko") // Add this line
	url_foto := c.FormValue("url_foto")
	id_foto_str := c.FormValue("id_foto")
	var photo_path string

	// Handle file upload
	file, err := c.FormFile("photo")
	if err == nil {
		// If file was uploaded successfully
		timestamp := time.Now().Format("20060102150405")
		filename := fmt.Sprintf("%s_%s", timestamp, file.Filename)
		photo_path = fmt.Sprintf("/uploads/%s", filename)

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
	}

	// Use provided URL or default if none provided
	if url_foto == "" {
		url_foto = "https://avatars.githubusercontent.com/u/174382151?s=48&v=4"
	}

	var id_foto uint = 1 // default value
	if id_foto_str != "" {
		if parsed, err := strconv.ParseUint(id_foto_str, 10, 32); err == nil {
			id_foto = uint(parsed)
		}
	}

	// Create photo URLs array
	photoURLs := []interface{}{
		map[string]string{
			"url":          url_foto,
			"originalName": file.Filename,
			"id_foto":      strconv.FormatUint(uint64(id_foto), 10), // Add id_foto to the map
		},
	}

	// Parse id_user if provided, otherwise use the token's user_id
	var input_user_id uint = uint(user_id)
	if id_user_str != "" {
		if parsed, err := strconv.ParseUint(id_user_str, 10, 32); err == nil {
			input_user_id = uint(parsed)
		}
	}

	input := models.StoreProcess{
		UserID:        input_user_id, // Use the parsed id_user
		NamaToko:      nama_toko,
		DeskripsiToko: deskripsi_toko, // Add this field
		URL:           url_foto,
		Photo:         photo_path, // Add photo path
		IdFoto:        id_foto,    // Now using uint type
		PhotoURLs:     photoURLs,  // Add this field to StoreProcess
	}

	response, err := handler.StoreService.Create(input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Handle photo upload with unique ID
	if len(photoURLs) > 0 {
		for _, photo := range photoURLs {
			photoModel := models.StorePhotoRequest{
				IdToko: response.ID,
				URL:    photo.(map[string]string)["url"],
				Photo:  photo.(map[string]string)["originalName"],
				// Let the database generate the ID
			}

			_, err := handler.StoreService.CreatePhoto(photoModel)
			if err != nil {
				return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
					Status:  false,
					Message: "Failed to create store photo",
					Error:   exceptions.NewString(err.Error()),
					Data:    nil,
				})
			}
		}
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to POST data",
		Error:   nil,
		Data:    response,
	})
}

func (handler *StoreHandler) DeleteStore(c *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to DELETE data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	user_id := claims.UserId
	id_toko, err := c.ParamsInt("id_toko")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to DELETE data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.StoreService.Delete(uint(id_toko), uint(user_id))
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
		Data:    response,
	})
}

type StoreInput struct {
	UserID   uint   `json:"user_id"`
	NamaToko string `json:"nama_toko"`
}
