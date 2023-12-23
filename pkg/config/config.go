package config

import (
	"log"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DBHOST"`
	DBName     string `mapstructure:"DBNAME"`
	DBUser     string `mapstructure:"DBUSER"`
	DBPort     string `mapstructure:"DBPORT"`
	DBPassword string `mapstructure:"DBPASSWORD"`
	PORT       string `mapstructure:"PORT"`
	SECRET_KEY string `mapstructure:"SECRET_KEY"`
}

var envs = []string{
	"DBHOST", "DBNAME", "DBUSER", "DBPORT", "DBPASSWORD", "PORT", "SECRET_KEY",
}

func LoadConfig() (Config, error) {
	var cfg Config
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return cfg, err
		}
	}
	cfgerr := viper.Unmarshal(&cfg)

	if err := validator.New().Struct(&cfg); err != nil {
		return cfg, err
	}
	LoadEnv()

	return cfg, cfgerr

}
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading Env File")
	}
}
