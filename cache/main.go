package cache

import (
	"ace/model"
	"go.uber.org/zap"
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
		&model.Type{},
	)

	if err != nil {
		zap.L().Error("[Mysql] AutoMigrate mysql table failure", zap.Error(err))
		os.Exit(1)
	}
}
