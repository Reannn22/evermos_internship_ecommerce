package utils

import "os"

func InitializeDirectories() error {
	dirs := []string{
		"uploads",
		"uploads/products",
		"uploads/stores",
		"uploads/users",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}
