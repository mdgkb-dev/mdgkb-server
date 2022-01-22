package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`

	DbDb       string `mapstructure:"DB_DB"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`

	RedisHost string `mapstructure:"REDIS_HOST"`
	RedisPort string `mapstructure:"REDIS_PORT"`

	UploadPath    string `mapstructure:"UPLOAD_PATH"`
	TemplatesPath string `mapstructure:"TEMPLATES_PATH"`

	TokenSecret string `mapstructure:"TOKEN_SECRET"`

	Email  Email  `mapstructure:",squash"`
	Social Social `mapstructure:",squash"`
}

type Email struct {
	User     string `mapstructure:"EMAIL_USER"`
	Password string `mapstructure:"EMAIL_PASSWORD"`
	From     string `mapstructure:"EMAIL_FROM"`
	Server   string `mapstructure:"EMAIL_SERVER"`
	Port     string `mapstructure:"EMAIL_PORT"`
}

type Social struct {
	InstagramToken string `mapstructure:"INSTAGRAM_TOKEN"`
	InstagramID    string `mapstructure:"INSTAGRAM_ID"`
}

func LoadConfig() (config *Config, err error) {
	//var result map[string]interface{}

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return config, err
}
