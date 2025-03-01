package handlers

import (
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

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService *services.UserService) UserHandler {
	return UserHandler{*userService}
}

func (handler *UserHandler) Route(app *fiber.App) {
	routes := app.Group("/api/v1/user")
	routes.Get("/", middleware.JWTProtected(), handler.GetAllUsers)
	routes.Get("/:id", middleware.JWTProtected(), handler.GetUserByID)
	routes.Put("/:id", middleware.JWTProtected(), handler.UserUpdate)
	routes.Delete("/:id", middleware.JWTProtected(), handler.UserDelete)
	routes.Get("/debug/count", middleware.JWTProtected(), handler.GetUserCount)
	routes.Patch("/:id/password", middleware.JWTProtected(), handler.ChangePassword)
	routes.Post("/forgot-password", handler.ForgotPassword)
	routes.Post("/reset-password", handler.ResetPassword)
}

func (handler *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	_, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	users, err := handler.userService.GetAllUsers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
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
		Data:    users,
	})
}

func (handler *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Get claims from JWT token
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Check if user is admin or accessing their own data
	if !claims.IsAdmin && uint(claims.UserId) != uint(id) {
		return c.Status(http.StatusForbidden).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Access denied: You can only view your own data",
			Error:   exceptions.NewString("forbidden access"),
			Data:    nil,
		})
	}

	user, err := handler.userService.GetUserByID(uint(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
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
		Data:    user,
	})
}

func (handler *UserHandler) UserUpdate(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Get claims from JWT token
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Check if user is admin or updating their own data
	if !claims.IsAdmin && uint(claims.UserId) != uint(id) {
		return c.Status(http.StatusForbidden).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Access denied: You can only update your own data",
			Error:   exceptions.NewString("forbidden access"),
			Data:    nil,
		})
	}

	var input models.UserRequest
	err = c.BodyParser(&input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to parse request data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses, err := handler.userService.Edit(uint(id), input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to update user",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Succeed to update user",
		Error:   nil,
		Data:    responses,
	})
}

func (handler *UserHandler) UserDelete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	_, err = jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.userService.Delete(uint(id))
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

func (handler *UserHandler) GetUserCount(c *fiber.Ctx) error {
	users, err := handler.userService.GetAllUsers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to count users",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "User count",
		Error:   nil,
		Data:    map[string]int{"count": len(users)},
	})
}

func (handler *UserHandler) ChangePassword(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID format",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	_, err = jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	var input models.PasswordChangeRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to parse request data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.userService.ChangePassword(uint(id), input.OldPassword, input.NewPassword)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to change password",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Password changed successfully",
		Error:   nil,
		Data:    response,
	})
}

func (handler *UserHandler) ForgotPassword(c *fiber.Ctx) error {
	var input models.ForgotPasswordRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to parse request data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.userService.ForgotPassword(input.Email)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Email not found",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	// Here you would typically send an email with the reset token
	// For this example, we'll just return the token in the response

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Password reset token generated successfully",
		Error:   nil,
		Data:    response,
	})
}

func (handler *UserHandler) ResetPassword(c *fiber.Ctx) error {
	var input models.ResetPasswordRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to parse request data",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.userService.ResetPassword(input.ResetToken, input.NewPassword)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to reset password",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Password reset successfully",
		Error:   nil,
		Data:    response,
	})
}
