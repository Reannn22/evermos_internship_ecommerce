package jwt

import "github.com/gofiber/fiber/v2"

// ...existing imports and code...

func ValidateToken(c *fiber.Ctx) error {
	// Try to extract token metadata
	_, err := ExtractTokenMetadata(c)
	if err != nil {
		return err
	}
	return nil
}
