package cache

import (
	"ace/model"
	"fmt"
	"os"
)

func migration() {
	err := DB.AutoMigrate(&model.User{})

	if err != nil {
		fmt.Printf("AutoMigrate error: %v", err.Error())
		os.Exit(1)
	}
}
