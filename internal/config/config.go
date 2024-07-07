package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	Application struct {
		Name string `mapstructure:"name"`
	} `mapstructure:"application"`

	Database struct {
		InfluxDB struct {
			URL         string `mapstructure:"url"`
			Token       string `mapstructure:"token"`
			Org         string `mapstructure:"org"`
			Bucket      string `mapstructure:"bucket"`
			Measurement struct {
				Gold     string `mapstructure:"gold"`
				Currency string `mapstructure:"currency"`
			} `mapstructure:"measurement"`
		} `mapstructure:"influxdb"`
	} `mapstructure:"database"`

	Rest struct {
		Altinkaynak struct {
			Gold     string `mapstructure:"gold"`
			Currency string `mapstructure:"currency"`
		} `mapstructure:"altinkaynak"`
	} `mapstructure:"rest"`

	HTTP struct {
		Client struct {
			Timeout int `mapstructure:"timeout"`
		} `mapstructure:"client"`
	} `mapstructure:"http"`

	Scheduler struct {
		Currency struct {
			Interval string `mapstructure:"interval"`
		} `mapstructure:"currency"`
	} `mapstructure:"scheduler"`
}

var (
	cfg  *Config
	once sync.Once
)

func LoadConfig() (*Config, error) {
	var err error
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("config")

		if err = viper.ReadInConfig(); err != nil {
			err = fmt.Errorf("config file error: %w", err)
			return
		}

		cfg = &Config{}
		if err = viper.Unmarshal(cfg); err != nil {
			err = fmt.Errorf("unable to decode into struct: %w", err)
		}
	})
	return cfg, err
}

func GetConfig() *Config {
	return cfg
}
