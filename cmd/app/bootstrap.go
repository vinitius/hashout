package app

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"viniti.us/hashout/config/log"
)

func Bootstrap(from string) {
	loadConfig(from)
	log.SetupLogger()
}

func loadConfig(from string) {
	env := getEnvWithDefault("env", "local")
	viper.AddConfigPath(from + "etc/")
	viper.SetConfigName("hashout-" + env)
	err := viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error %w", err))
	}

	fmt.Printf("[hashout] Config loaded from %s", viper.ConfigFileUsed())
}

func getEnvWithDefault(envVar, defaultValue string) string {
	if value, ok := os.LookupEnv(envVar); ok {
		if value != "" {
			return value
		}

	}
	return defaultValue
}

func Run() error {
	log.Logger.Info("Starting hashout...")

	app := SetupApplication()
	return app.ListenAndServe()
}
