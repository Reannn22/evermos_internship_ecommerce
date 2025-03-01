package handlers

import (
	"fmt"
	"mini-project-evermos/exceptions"
	"mini-project-evermos/models"
	"mini-project-evermos/models/responder"
	"mini-project-evermos/services"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TransactionDetailHandler struct {
	service services.TransactionDetailService
}

func NewTransactionDetailHandler(service services.TransactionDetailService) *TransactionDetailHandler {
	return &TransactionDetailHandler{
		service: service,
	}
}

func (handler *TransactionDetailHandler) Route(app *fiber.App) {
	trxDetail := app.Group("/api/v1/detail-trx")

	trxDetail.Get("/", handler.GetAll)
	trxDetail.Get("/:id", handler.GetById)
	trxDetail.Get("/transaction/:trxId", handler.GetByTrxId)
	trxDetail.Post("/", handler.Create)
	trxDetail.Put("/:id", handler.Update)
	trxDetail.Delete("/:id", handler.Delete)
}

func (handler *TransactionDetailHandler) GetAll(c *fiber.Ctx) error {
	details, err := handler.service.GetAll()
	if err != nil {
		fmt.Printf("Error in GetAll handler: %v\n", err) // Add logging
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "INTERNAL SERVER ERROR",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Check if details is nil or empty
	if details == nil {
		details = []models.TransactionDetailResponse{} // Return empty array instead of null
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Error:   nil,
		Data:    details,
	})
}

func (handler *TransactionDetailHandler) GetById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	detail, err := handler.service.GetById(uint(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Transaction detail not found",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Error:   nil,
		Data:    detail,
	})
}

func (handler *TransactionDetailHandler) GetByTrxId(c *fiber.Ctx) error {
	trxId, err := strconv.Atoi(c.Params("trxId"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid transaction ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	details, err := handler.service.GetByTrxId(uint(trxId))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Transaction details not found",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Error:   nil,
		Data:    details,
	})
}

func (handler *TransactionDetailHandler) Create(c *fiber.Ctx) error {
	var input models.TransactionDetailProcess
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid input",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	result, err := handler.service.Create(input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to POST data",
		Error:   nil,
		Data:    result,
	})
}

func (handler *TransactionDetailHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	var input models.TransactionDetailProcess
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid input",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	result, err := handler.service.Update(uint(id), input)
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
		Data:    result,
	})
}

func (handler *TransactionDetailHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	detail, err := handler.service.Delete(uint(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
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
		Data:    detail,
	})
}
