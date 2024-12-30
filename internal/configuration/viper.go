package configuration

import (
	"errors"
	"flag"
	"log/slog"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {

	// create new viper instance
	Config = viper.New()

	flag.Bool("verbose", false, "enable verbose logging")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := Config.BindPFlags(pflag.CommandLine)
	if err != nil {
		slog.Error("unable to bind flags for configuration", "error", err.Error())
	}

	if Config.GetBool("verbose") {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	// setup default values for the http server
	Config.SetDefault("http.host", "0.0.0.0")
	Config.SetDefault("http.port", 8000)
	Config.SetDefault("http.trusted_proxies", []string{"localhost", "127.0.0.1"})

	// setup the configuration loading process
	Config.SetConfigName("config")
	Config.AddConfigPath(".")
	Config.AddConfigPath("/etc/gegenlicht-register/")

	// read the configuration
	if err := Config.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			slog.Error("no configuration file found")
			os.Exit(1)
		}
		slog.Error("unable to read configuration file", "error", err.Error())
	}
}
