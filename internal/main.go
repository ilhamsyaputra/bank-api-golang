package main

import (
	"bank-api/internal/core/services"
	"bank-api/internal/handlers"
	"bank-api/internal/repositories"
	"bank-api/internal/server"
	"bank-api/pkg/logger"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// Service Name
	SERVICE := viper.GetString("SERVICE")

	// Database configuration
	DB_DRIVER := viper.GetString("DB_DRIVER")
	DB_USER := viper.GetString("DB_USER")
	DB_PASSWORD := viper.GetString("DB_PASSWORD")
	DB_HOST := viper.GetString("DB_HOST")
	DB_PORT := viper.GetInt("DB_PORT")
	DB_DATABASE := viper.GetString("DB_DATABASE")

	// SERVICE configuration
	// SERVICE_HOST := viper.GetString("SERVICE_HOST")
	// SERVICE_PORT := viper.GetInt("SERVICE_HOST")

	// Dependency injection
	logger := logger.NewLogger(SERVICE)
	repository := repositories.InitRepository(DB_DRIVER, DB_HOST, DB_USER, DB_PASSWORD, DB_DATABASE, DB_PORT, logger)
	service := services.InitBankService(repository, logger)
	handler := handlers.InitHandler(service, logger)
	server := server.InitServer(handler)

	// Start service API
	server.Start()
}
