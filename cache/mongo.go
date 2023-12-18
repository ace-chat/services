package cache

import (
	"ace/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"os"
)

var Mongo *mongo.Database

func InitMongo(conf model.Mongo) {
	ctx := context.Background()
	var err error
	clientOptions := options.Client().ApplyURI(conf.Uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		zap.L().Error("[Mongo] Connect mongodb failed", zap.Error(err))
		os.Exit(0)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		zap.L().Error("[Mongo] Ping mongodb failed", zap.Error(err))
		os.Exit(0)
	}

	Mongo = client.Database(conf.DB)
}
