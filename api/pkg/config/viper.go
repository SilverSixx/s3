package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	storageUrl  string `json:"storage"`
}

func LoadConfig() (*Config, error) {

	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AddConfigPath(".")

	v.AutomaticEnv()

	setDefaultValues(v)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("No configuration file found.")
		} else {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error unmarshalling config, %s", err)
		return nil, err
	}

	viper.MergeConfigMap(v.AllSettings())

	return &cfg, nil
}

func setDefaultValues(v *viper.Viper) {
	v.SetDefault("storageUrl", "https://s3.amazonaws.com")
}