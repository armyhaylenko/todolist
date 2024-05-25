package logging

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func SetupLogger(isDevelopment bool) {
	var config zap.Config
	if isDevelopment {
		config = setupDevelopmentConfig()
	} else {
		config = setupProductionConfig()
	}
	logger := zap.Must(config.Build())

	defer logger.Sync()
	Logger = logger.Sugar()
}

// Creates initial development config for the logger.
// Customize further if needed.
func setupDevelopmentConfig() zap.Config {
	config := zap.NewDevelopmentConfig()

	return config
}

// Creates initial production config for the logger.
// Customize further if needed.
func setupProductionConfig() zap.Config {
	config := zap.NewProductionConfig()

	return config
}
