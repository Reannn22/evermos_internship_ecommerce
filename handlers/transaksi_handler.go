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

type TransactionHandler struct {
	TransactionService services.TransactionService
}

func NewTransactionHandler(transactionService *services.TransactionService) TransactionHandler {
	return TransactionHandler{*transactionService}
}

func (handler *TransactionHandler) Route(app *fiber.App) {
	routes := app.Group("/api/v1/trx")

	// Public routes (with JWT protection only)
	routes.Get("/", middleware.JWTProtected(), handler.GetAllTransaction)
	routes.Get("/:id", middleware.JWTProtected(), handler.DetailTransaction)

	// Admin only routes
	routes.Post("/", middleware.JWTProtected(), middleware.RolePermissionAdmin(), handler.CreateTransaction)
	routes.Put("/:id", middleware.JWTProtected(), middleware.RolePermissionAdmin(), handler.UpdateTransaction)
	routes.Delete("/:id", middleware.JWTProtected(), middleware.RolePermissionAdmin(), handler.DeleteTransaction)
}

func cleanTransactionResponse(response models.TransactionResponse) models.TransactionResponse {
	// If there are no transaction details, return response without that field
	if len(response.TransactionDetails) == 0 {
		return models.TransactionResponse{
			ID:          response.ID,
			UserID:      response.UserID, // Make sure to include UserID
			HargaTotal:  response.HargaTotal,
			KodeInvoice: response.KodeInvoice,
			MethodBayar: response.MethodBayar,
			Address:     response.Address,
			CreatedAt:   response.CreatedAt,
			UpdatedAt:   response.UpdatedAt,
			// TransactionDetails is omitted intentionally
		}
	}
	return response
}

func (handler *TransactionHandler) GetAllTransaction(c *fiber.Ctx) error {
	// Default values
	defaultLimit := 10
	defaultPage := 1

	limit, err := strconv.Atoi(c.FormValue("limit", strconv.Itoa(defaultLimit)))
	if err != nil {
		limit = defaultLimit
	}

	page, err := strconv.Atoi(c.FormValue("page", strconv.Itoa(defaultPage)))
	if err != nil {
		page = defaultPage
	}

	keyword := c.FormValue("search", "")

	responses, err := handler.TransactionService.GetAll(limit, page, keyword)
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

func (handler *TransactionHandler) DetailTransaction(c *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
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

	response, err := handler.TransactionService.GetById(uint(id), uint(claims.UserId))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
				Status:  false,
				Message: "Transaction not found",
				Error:   exceptions.NewString(err.Error()),
				Data:    nil,
			})
		}
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
		Data:    cleanTransactionResponse(response),
	})
}

func (handler *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		fmt.Printf("JWT Error: %v\n", err)
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}
	fmt.Printf("User ID from token: %d\n", claims.UserId)

	var input models.TransactionRequest
	err = c.BodyParser(&input)
	if err != nil {
		fmt.Printf("Body Parse Error: %v\n", err)
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to parse request body",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Add validation for alamat_pengiriman
	if input.AlamatPengiriman == 0 {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Alamat diperlukan untuk alamat kirim produk",
			Error:   exceptions.NewString("alamat_pengiriman is required"),
			Data:    nil,
		})
	}

	// Use the user_id from the request body instead of claims
	fmt.Printf("Request user_id: %d\n", input.UserID)

	response, err := handler.TransactionService.Create(input, input.UserID) // Use input.UserID instead of claims.UserId
	if err != nil {
		fmt.Printf("Service Error: %v\n", err)
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Status:  false,
			Message: "NOT FOUND",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to POST data",
		Error:   nil,
		Data:    cleanTransactionResponse(response),
	})
}

func (handler *TransactionHandler) UpdateTransaction(c *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID parameter",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	var input models.TransactionUpdateRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to parse request body",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Add validation for alamat_pengiriman
	if input.AlamatPengiriman == 0 {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Alamat diperlukan untuk alamat kirim produk",
			Error:   exceptions.NewString("alamat_pengiriman is required"),
			Data:    nil,
		})
	}

	response, err := handler.TransactionService.Update(uint(id), uint(claims.UserId), input)
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
				Status:  false,
				Message: "Transaction not found",
				Error:   exceptions.NewString(err.Error()),
				Data:    nil,
			})
		}
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to update transaction",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Transaction updated successfully",
		Error:   nil,
		Data:    cleanTransactionResponse(response),
	})
}

func (handler *TransactionHandler) DeleteTransaction(c *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Get transaction data before deletion
	transaction, err := handler.TransactionService.GetById(uint(id), uint(claims.UserId))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get transaction",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Delete the transaction
	err = handler.TransactionService.Delete(uint(id), uint(claims.UserId))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to delete transaction",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Return the deleted transaction data
	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully deleted transaction",
		Error:   nil,
		Data:    cleanTransactionResponse(transaction),
	})
}
