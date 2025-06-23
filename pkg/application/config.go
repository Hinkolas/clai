package application

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func LoadConfig(path string, name string) (*Config, error) {

	v := viper.NewWithOptions(viper.KeyDelimiter("|"))

	// Define default config values
	// v.SetDefault("logging|log_level", "debug")

	// Tell viper where to look for the config file
	v.SetConfigName(name)
	v.AddConfigPath(path)
	v.SetConfigType("yaml")

	fmt.Println(filepath.Join(path, name))

	// Check if config file exists
	if _, err := os.Stat(filepath.Join(path, name)); os.IsNotExist(err) {
		// File doesn't exist, create it with defaults
		if err := v.SafeWriteConfig(); err != nil {
			return nil, fmt.Errorf("failed to create config file: %w", err)
		}
	}

	// Read the config
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	// Unmarshal the config into application.Config
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}

	return &cfg, nil

}

type Config struct {
	//Models  map[string]llm.Model `mapstructure:"models" yaml:"models"`
}
