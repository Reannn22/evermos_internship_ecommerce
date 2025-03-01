package handlers

import (
	"mini-project-evermos/exceptions"
	"mini-project-evermos/middleware"
	"mini-project-evermos/models"
	"mini-project-evermos/models/responder"
	"mini-project-evermos/services"
	"mini-project-evermos/utils/jwt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductHandler struct {
	ProductService services.ProductService
}

func NewProductHandler(productService *services.ProductService) ProductHandler {
	return ProductHandler{*productService}
}

func (handler *ProductHandler) Route(app *fiber.App) {
	routes := app.Group("/api/v1/produk")

	// Public routes (with JWT protection only)
	routes.Get("/", middleware.JWTProtected(), handler.GetAllProduct)
	routes.Get("/search", middleware.JWTProtected(), handler.SearchProducts)
	routes.Get("/:id", middleware.JWTProtected(), handler.ProductDetail)
	routes.Get("/category/:category", middleware.JWTProtected(), handler.GetProductsByCategory)
	routes.Get("/:id/related", middleware.JWTProtected(), handler.GetRelatedProducts)
	routes.Get("/photo/:id", handler.ServeProductPhoto)

	// Admin only routes
	routes.Post("/", middleware.JWTProtected(), middleware.RolePermissionAdmin(), handler.ProductCreate)
	routes.Put("/:id", middleware.JWTProtected(), middleware.RolePermissionAdmin(), handler.ProductUpdate)
	routes.Delete("/:id", middleware.JWTProtected(), middleware.RolePermissionAdmin(), handler.ProductDelete)
}

func (handler *ProductHandler) GetAllProduct(c *fiber.Ctx) error {
	limit := 10
	page := 1
	keyword := ""

	if c.Query("limit") != "" {
		if val, err := strconv.Atoi(c.Query("limit")); err == nil {
			limit = val
		}
	}

	if c.Query("page") != "" {
		if val, err := strconv.Atoi(c.Query("page")); err == nil {
			page = val
		}
	}

	if q := c.Query("q"); q != "" {
		keyword = q
	}

	responses, err := handler.ProductService.FindAllPagination(limit, page, keyword)
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

func (handler *ProductHandler) SearchProducts(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Search query is required",
			Error:   exceptions.NewString("query parameter 'q' is missing"),
			Data:    nil,
		})
	}

	products, err := handler.ProductService.SearchProducts(query)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to search products",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to search products",
		Error:   nil,
		Data:    products,
	})
}

func (handler *ProductHandler) ProductDetail(c *fiber.Ctx) error {
	_, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.ProductService.FindById(uint(id))
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

func (handler *ProductHandler) ProductCreate(c *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Parse store_id from form data
	store_id, err := strconv.Atoi(c.FormValue("store_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid store_id",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to parse form",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	category_id, err := strconv.Atoi(c.FormValue("category_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid category_id",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	stok, err := strconv.Atoi(c.FormValue("stok"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid stok",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	files := form.File["photo_files"]
	var photoURLs []interface{}

	// Handle file upload first
	if len(files) > 0 {
		file := files[0] // Take only the first file
		ext := filepath.Ext(file.Filename)
		newFilename := uuid.New().String() + ext
		filename := filepath.Join("uploads", "products", newFilename)

		if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
				Status:  false,
				Message: "Failed to create upload directory",
				Error:   exceptions.NewString(err.Error()),
				Data:    nil,
			})
		}

		if err := c.SaveFile(file, filename); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
				Status:  false,
				Message: "Failed to save file",
				Error:   exceptions.NewString(err.Error()),
				Data:    nil,
			})
		}

		fileURL := "/" + strings.ReplaceAll(filename, "\\", "/")
		originalURL := c.FormValue("photo_urls") // Get the original URL

		// Create a combined photo entry
		photoURLs = append(photoURLs, map[string]string{
			"url":          originalURL, // This will be mapped to "url" in response
			"originalName": fileURL,     // This will be mapped to "photo" in response
		})
	}

	input := models.ProductRequest{
		NamaProduk:    c.FormValue("nama_produk"),
		CategoryID:    uint(category_id),
		StoreID:       uint(store_id), // Use the parsed store_id from form
		HargaReseller: c.FormValue("harga_reseller"),
		HargaKonsumen: c.FormValue("harga_konsumen"),
		Stok:          stok,
		Deskripsi:     c.FormValue("deskripsi"),
		PhotoURLs:     photoURLs,
	}

	// Pass the user ID separately for authorization
	response, err := handler.ProductService.Create(input, uint(claims.UserId))
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

func (handler *ProductHandler) ProductUpdate(c *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to PUT data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to PUT data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	category_id, err := strconv.Atoi(c.FormValue("category_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid category_id",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	store_id, err := strconv.Atoi(c.FormValue("store_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid store_id",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	stok, err := strconv.Atoi(c.FormValue("stok"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid stok",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	var photoID uint
	if photoIDStr := c.FormValue("photo_id"); photoIDStr != "" {
		if val, err := strconv.ParseUint(photoIDStr, 10, 32); err == nil {
			photoID = uint(val)
		}
	}

	var photoURLs []interface{}

	// Handle file upload first
	if form, err := c.MultipartForm(); err == nil {
		if files := form.File["photo_files"]; len(files) > 0 {
			file := files[0] // Take only the first file
			ext := filepath.Ext(file.Filename)
			newFilename := uuid.New().String() + ext
			filename := filepath.Join("uploads", "products", newFilename)

			if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
				return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
					Status:  false,
					Message: "Failed to create upload directory",
					Error:   exceptions.NewString(err.Error()),
					Data:    nil,
				})
			}

			if err := c.SaveFile(file, filename); err != nil {
				return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
					Status:  false,
					Message: "Failed to save file",
					Error:   exceptions.NewString(err.Error()),
					Data:    nil,
				})
			}

			fileURL := "/" + strings.ReplaceAll(filename, "\\", "/")
			originalURL := c.FormValue("photo_urls") // Get the original URL

			// Create a combined photo entry
			photoURLs = append(photoURLs, map[string]string{
				"url":          originalURL, // This will be mapped to "url" in response
				"originalName": fileURL,     // This will be mapped to "photo" in response
			})
		} else if photoURL := c.FormValue("photo_urls"); photoURL != "" {
			photoURLs = append(photoURLs, map[string]string{
				"url":          photoURL,                   // This will be mapped to "url" in response
				"originalName": c.FormValue("photo_files"), // This will be mapped to "photo" in response
			})
		}
	}

	input := models.ProductRequest{
		NamaProduk:    c.FormValue("nama_produk"),
		CategoryID:    uint(category_id),
		StoreID:       uint(store_id),
		HargaReseller: c.FormValue("harga_reseller"),
		HargaKonsumen: c.FormValue("harga_konsumen"),
		Stok:          stok,
		Deskripsi:     c.FormValue("deskripsi"),
		PhotoURLs:     photoURLs,
		PhotoIDs:      []uint{photoID},
	}

	response, err := handler.ProductService.Update(uint(id), input, uint(claims.UserId))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
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

func (handler *ProductHandler) ProductDelete(c *fiber.Ctx) error {
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

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to DELETE data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	deletedProduct, err := handler.ProductService.Delete(uint(id), uint(user_id))
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
		Data:    deletedProduct,
	})
}

func (handler *ProductHandler) GetProductsByCategory(c *fiber.Ctx) error {
	_, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	categoryID := c.Params("category")
	if categoryID == "" {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Category ID is required",
			Error:   exceptions.NewString("category parameter is missing"),
			Data:    nil,
		})
	}

	products, err := handler.ProductService.FindByCategory(categoryID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get products by category",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to GET products by category",
		Error:   nil,
		Data:    products,
	})
}

func (handler *ProductHandler) GetRelatedProducts(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	relatedProducts, err := handler.ProductService.GetRelatedProducts(uint(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get related products",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully retrieved related products",
		Error:   nil,
		Data:    relatedProducts,
	})
}

func (handler *ProductHandler) ServeProductPhoto(c *fiber.Ctx) error {
	photoID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid photo ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	photo, err := handler.ProductService.GetProductPhoto(uint(photoID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Photo not found",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Photo found",
		Error:   nil,
		Data:    photo,
	})
}
