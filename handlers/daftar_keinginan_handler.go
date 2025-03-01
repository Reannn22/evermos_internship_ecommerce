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

type WishlistHandler struct {
	WishlistService services.WishlistService
}

func NewWishlistHandler(wishlistService *services.WishlistService) WishlistHandler {
	return WishlistHandler{*wishlistService}
}

func (handler *WishlistHandler) Route(app *fiber.App) {
	routes := app.Group("/api/v1/wishlist-shopping")
	routes.Use(middleware.JWTProtected())

	// Put the /clear route before any routes with parameters
	routes.Delete("/clear", handler.ClearAll)

	routes.Get("/", handler.GetAll)
	routes.Get("/:id", handler.GetById)
	routes.Post("/", handler.Create)
	routes.Put("/:id", handler.Update)
	routes.Delete("/:id", handler.Delete)
}

func (handler *WishlistHandler) GetAll(c *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	// Add debug print
	fmt.Printf("Getting wishlists for user ID: %d\n", claims.UserId)

	responses, err := handler.WishlistService.GetAll(uint(claims.UserId))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get wishlists",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success get wishlists",
		Data:    responses,
	})
}

func (handler *WishlistHandler) GetById(c *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	response, err := handler.WishlistService.GetById(uint(id), uint(claims.UserId))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get wishlist",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success get wishlist",
		Data:    response,
	})
}

func (handler *WishlistHandler) Create(c *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	var input models.WishlistRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid request body",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	// Parse store_id from form if not in JSON body
	if input.StoreID == 0 {
		storeID, err := strconv.ParseUint(c.FormValue("store_id"), 10, 64)
		if err == nil {
			input.StoreID = uint(storeID)
		}
	}

	response, err := handler.WishlistService.Create(input, uint(claims.UserId))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to create wishlist",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success create wishlist",
		Data:    response,
	})
}

func (handler *WishlistHandler) Update(c *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	var input models.WishlistRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid request body",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	response, err := handler.WishlistService.Update(uint(id), input, uint(claims.UserId))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to update wishlist",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success update wishlist",
		Data:    response,
	})
}

func (handler *WishlistHandler) Delete(c *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Invalid ID",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	// Get wishlist data before deletion
	wishlist, err := handler.WishlistService.GetById(uint(id), uint(claims.UserId))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to get wishlist",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	// Delete the wishlist
	err = handler.WishlistService.Delete(uint(id), uint(claims.UserId))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to delete wishlist",
			Error:   exceptions.NewString(err.Error()),
		})
	}

	return c.JSON(responder.ApiResponse{
		Status:  true,
		Message: "Success delete wishlist",
		Data:    wishlist, // Return the wishlist data that was just deleted
	})
}

func (handler *WishlistHandler) ClearAll(c *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	deletedWishlists, err := handler.WishlistService.ClearAll(uint(claims.UserId))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Status:  false,
			Message: "Failed to clear wishlist",
			Error:   exceptions.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Status:  true,
		Message: "Successfully cleared wishlist",
		Error:   nil,
		Data:    deletedWishlists,
	})
}
