package base

import "github.com/spf13/viper"

var cfg *Config

type Config struct {
	Mode    string
	Port    int
	Version string
	Domain  string
}

// Init initial base configuration
func Init() {
	mode := viper.GetString("mode")
	if mode == "" {
		mode = "dev"
	}

	port := viper.GetInt("port")
	if port <= 0 {
		port = 8080
	}
	version := viper.GetString("version")
	if version == "" {
		version = "未设置版本号"
	}
	domain := viper.GetString("domain")
	if domain == "" {
		domain = "127.0.0.1"
	}

	cfg = &Config{
		Mode:    mode,
		Port:    port,
		Version: version,
		Domain:  domain,
	}
}

// GetMode return the application running mode
func GetMode() string {
	return cfg.Mode
}

// GetPort return the port to which the application is bound
func GetPort() int {
	return cfg.Port
}

// GetVersion return the version of the application
func GetVersion() string {
	return cfg.Version
}

func GetDomain() string {
	return cfg.Domain
}
