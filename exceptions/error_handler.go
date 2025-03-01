package exceptions

import (
	"mini-project-evermos/models/responder"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewString(s string) *string {
	return &s
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)

	if ok {
		return ctx.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			// Status: "BAD_REQUEST",
			Status:  false,
			Message: "BAD REQUEST",
			Error:   NewString(err.Error()),
			Data:    nil,
		})
	}

	// exception, ok := err.(NotFoundError)

	// if ok {
	// 	return ctx.Status(http.StatusNotFound).JSON(responder.ApiResponse{
	// 		// Status: "NOT_FOUND",
	// 		Code:    http.StatusNotFound,
	// 		Message: "NOT FOUND",
	// 		Error:   &exception.Error,
	// 		Data:    nil,
	// 	})
	// }

	return ctx.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
		// Status: "INTERNAL_SERVER_ERROR",
		Status:  false,
		Message: "INTERNAL SERVER ERROR",
		Error:   NewString(err.Error()),
		Data:    nil,
	})
}
