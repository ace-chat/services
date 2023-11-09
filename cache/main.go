package cache

import (
	"ace/model"
	"fmt"
	"os"
)

func migration() {
	err := DB.AutoMigrate(
		&model.User{},
		&model.Plan{},
		&model.UserPlan{},
		&model.MediaAds{},
		&model.MediaContent{},
		&model.EmailAds{},
		&model.EmailContent{},
		&model.OptimizedContent{},
		&model.OptimizedAds{},
		&model.EngineAds{},
		&model.EngineContent{},
		&model.BlogAds{},
		&model.BlogContent{},
		&model.Tone{},
		&model.Voice{},
		&model.Gender{},
		&model.Region{},
		&model.Platform{},
		&model.Language{},
	)

	if err != nil {
		fmt.Printf("AutoMigrate error: %v", err.Error())
		os.Exit(1)
	}
}
