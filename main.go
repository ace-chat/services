package main

import (
	"ace/cache"
	"ace/conf"
	"ace/server"
	"flag"
	"fmt"
	"os"
)

func main() {
	var path string
	flag.StringVar(&path, "c", "./conf/conf.yaml", "configuration file path")
	flag.Parse()
	config := conf.GetConf(path)

	cache.InitRedis(config.Redis)
	cache.InitMysql(config.Mysql)

	r := server.NewServer(config.Mode)

	err := r.Run(fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		fmt.Printf("Start server failed, error: %v \n", err.Error())
		os.Exit(0)
	}
}
