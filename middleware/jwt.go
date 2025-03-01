package middleware

import (
	"mini-project-evermos/exceptions"
	"mini-project-evermos/models/responder"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
)

// JWTProtected func for specify routes group with JWT authentication.
func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	return c.Status(http.StatusUnauthorized).JSON(responder.ApiResponse{
		Status:  false,
		Message: "Something Wrong",
		Error:   exceptions.NewString("Unauthorized"),
		Data:    nil,
	})
}
