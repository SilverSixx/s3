package main

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"

	v "github.com/silversixx/s3-go/pkg/config"
	"github.com/silversixx/s3-go/pkg/logger"
	api "github.com/silversixx/s3-go/pkg/server"
)

func main() {
	v.LoadConfig()
	logger.Initialize(viper.GetString("log_format"), zap.InfoLevel)
	server := &api.ApiServer{}
	server.InitServer()
	server.Run()
}
