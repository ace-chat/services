package main

import (
	"ace/cache"
	"ace/conf"
	"ace/logger"
	"ace/pkg"
	"ace/request"
	"ace/server"
	"flag"
	"fmt"
	"os"

	"go.uber.org/zap"
)

func main() {
	var path string
	flag.StringVar(&path, "c", "./conf/conf.yaml", "configuration file path")
	flag.Parse()
	config := conf.GetConf(path)

	logger.NewLogger(config.Mode, config.Name, config.Logger)
	request.Setup(config.Request)

	cache.InitRedis(config.Redis)
	cache.InitMysql(config.Mysql)
	cache.InitMongo(config.Mongo)

	pkg.Init(config.Upload)

	r := server.NewServer(config.Mode)

	err := r.Run(fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		zap.L().Error("[Service] Start server failed", zap.Error(err))
		os.Exit(0)
	}
}
