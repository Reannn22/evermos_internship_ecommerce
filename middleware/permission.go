package middleware

import (
	"mini-project-evermos/exceptions"
	"mini-project-evermos/models/responder"
	"mini-project-evermos/utils/jwt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Modify RolePermissionAdmin to not require database parameter
func RolePermissionAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, err := jwt.ExtractTokenMetadata(c)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
				Status:  false,
				Message: "Unauthorized",
				Error:   exceptions.NewString(err.Error()),
				Data:    nil,
			})
		}

		// Just check IsAdmin from JWT claims
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
}
