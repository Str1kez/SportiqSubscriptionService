package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Server    *ServerConfig
	HistoryDB *HistoryDBConfig
	DB        *DBConfig
	MQ        *MQConfig
}

type ServerConfig struct {
	Host string `mapstructure:"APP_HOST"`
	Port string `mapstructure:"APP_PORT"`
}

type HistoryDBConfig struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	DBName   string `mapstructure:"POSTGRES_DB"`
}

type DBConfig struct {
	Url string `mapstructure:"DB_SUB_URL"`
}

type MQConfig struct {
	ConsumerCount uint8  `mapstructure:"CONSUMER_COUNT"`
	Host          string `mapstructure:"RABBITMQ_HOST"`
	Port          string `mapstructure:"RABBITMQ_PORT"`
	User          string `mapstructure:"RABBITMQ_DEFAULT_USER"`
	Password      string `mapstructure:"RABBITMQ_DEFAULT_PASS"`
}

// type ErrorConfig struct {
// 	Completion string `mapstructure:"completion"`
// 	Sending    string `mapstructure:"sending"`
// 	Converting string `mapstructure:"converting"`
// 	Parsing    string `mapstructure:"parsing"`
// 	Context    string `mapstructure:"context"`
// }

func NewConfig() (*Config, error) {
	var config Config

	err := parseEnv()
	if err != nil {
		log.Errorf("Error in parsing env data: %v\n", err)
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Errorf("Error in unmarshalling env data: %v\n", err)
		return nil, err
	}
	err = viper.Unmarshal(&config.Server)
	if err != nil {
		log.Errorf("Error in unmarshalling env data: %v\n", err)
		return nil, err
	}
	err = viper.Unmarshal(&config.DB)
	if err != nil {
		log.Errorf("Error in unmarshalling env data: %v\n", err)
		return nil, err
	}
	err = viper.Unmarshal(&config.MQ)
	if err != nil {
		log.Errorf("Error in unmarshalling env data: %v\n", err)
		return nil, err
	}
	err = viper.Unmarshal(&config.HistoryDB)
	if err != nil {
		log.Errorf("Error in unmarshalling env data: %v\n", err)
		return nil, err
	}

	return &config, nil
}

func parseEnv() error {
	viper.SetConfigFile(".dev.env")
	if err := viper.ReadInConfig(); err == nil {
		return nil
	}
	if err := viper.BindEnv("APP_PORT"); err != nil {
		return err
	}
	if err := viper.BindEnv("APP_HOST"); err != nil {
		return err
	}
	if err := viper.BindEnv("POSTGRES_DB"); err != nil {
		return err
	}
	if err := viper.BindEnv("POSTGRES_HOST"); err != nil {
		return err
	}
	if err := viper.BindEnv("POSTGRES_USER"); err != nil {
		return err
	}
	if err := viper.BindEnv("POSTGRES_PORT"); err != nil {
		return err
	}
	if err := viper.BindEnv("POSTGRES_PASSWORD"); err != nil {
		return err
	}
	if err := viper.BindEnv("DB_SUB_URL"); err != nil {
		return err
	}
	if err := viper.BindEnv("RABBITMQ_HOST"); err != nil {
		return err
	}
	if err := viper.BindEnv("RABBITMQ_PORT"); err != nil {
		return err
	}
	if err := viper.BindEnv("RABBITMQ_DEFAULT_USER"); err != nil {
		return err
	}
	if err := viper.BindEnv("RABBITMQ_DEFAULT_PASS"); err != nil {
		return err
	}
	if err := viper.BindEnv("CONSUMER_COUNT"); err != nil {
		return err
	}
	return nil
}
