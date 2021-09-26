package log

import (
	defaultLog "log"
	"strings"
	"sync"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once   = new(sync.Once)
	Logger *zap.SugaredLogger
)

func SetupLogger() {
	once.Do(func() {
		log, err := newLogger()
		if err != nil {
			defaultLog.Fatalf("Logger could not be initialized: %v", err)
		}
		Logger = log.Sugar()
		Logger.Debug("Logger setup has finished")
	})
}

func newLogger() (*zap.Logger, error) {
	env := viper.GetString("ENV")

	var config zap.Config
	if env == "local" {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	lvl := viper.GetString("LOG_LEVEL")

	if exists := lvl != ""; exists {
		lvl = strings.ToLower(lvl)
		switch lvl {
		case "debug":
			config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		case "info":
			config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		case "warn":
			config.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
		case "error":
			config.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
		case "panic":
			config.Level = zap.NewAtomicLevelAt(zapcore.PanicLevel)
		case "fatal":
			config.Level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
		}
	}

	return config.Build()
}
