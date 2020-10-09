package main

import (
	"log"
	"minio_example/cmd/internal/cloudstorage"
	"minio_example/cmd/internal/handlers"
	"minio_example/cmd/internal/settings"
	"net/http"

	"go.uber.org/zap"
)

func main() {
	appSettings, err := settings.New()
	if err != nil {
		log.Fatal(err)
	}
	logger, err := getConfiguredLogger(appSettings.Debug, appSettings.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	s3Client, err := cloudstorage.NewS3Client(logger, appSettings.AWSSettings)
	if err != nil {
		log.Fatal(err)
	}

	indexHandler := handlers.NewIndexHandler(logger, s3Client, appSettings.Bucket)
	http.Handle("/", indexHandler)
	http.ListenAndServe(":80", nil)
}

func getConfiguredLogger(debug bool, logLevel string) (logger *zap.Logger, err error) {
	var loggerConfig zap.Config
	if debug {
		loggerConfig = zap.NewDevelopmentConfig()
	} else {
		loggerConfig = zap.NewProductionConfig()
	}

	if err := loggerConfig.Level.UnmarshalText([]byte(logLevel)); err != nil {
		return nil, err
	}

	if logger, err = loggerConfig.Build(); err != nil {
		return nil, err
	}
	return logger, nil
}