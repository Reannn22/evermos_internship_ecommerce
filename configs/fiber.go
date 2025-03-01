package configs

import (
	"mini-project-evermos/exceptions"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exceptions.ErrorHandler,
	}
}
