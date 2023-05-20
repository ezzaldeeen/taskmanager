package middlewares

import (
	"go.uber.org/zap"
	"log"
	"taskmanager/config"
)

const (
	dev  string = "development"
	prod string = "production"
)

func GetLogging() func() {
	// Load app's configuration
	cfg := config.Load()
	switch cfg.Environment {
	case dev:
		return DevLogging
	case prod:
		return ProdLogging
	}
	return nil
}

func DevLogging() {
	_, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("unable to load logger", err)
	}
}

func ProdLogging() {
	_, err := zap.NewProduction()
	if err != nil {
		log.Fatal("unable to load logger", err)
	}
}
