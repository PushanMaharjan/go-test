package lib

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort         string `mapstructure:"SERVER_PORT"`
	Environment        string `mapstructure:"ENVIRONMENT"`
	DBUsername         string `mapstructure:"DB_USER"`
	DBPassword         string `mapstructure:"DB_PASS"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBName             string `mapstructure:"DB_NAME"`
	MaxMultipartMemory int64  `mapstructure:"MAX_MULTIPART_MEMORY"`
	MailClientId       string `mapstructure:"MAIL_CLIENT_ID"`
	MailClientSecret   string `mapstructure:"MAIL_CLIENT_SECRET"`
	MailRefreshToken   string `mapstructure:"MAIL_REFRESH_TOKEN"`
	MailAccessToken    string `mapstructure:"MAIL_ACCESS_TOKEN"`
}

var globalEnv = Env{
	MaxMultipartMemory: 10 << 20, // 10 MB
}

func GetEnv() Env {
	return globalEnv
}

func NewEnv() Env {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("cannot read cofiguration")
	}

	err = viper.Unmarshal(&globalEnv)
	if err != nil {
		log.Println("environment cant be loaded: ", err)
	}

	return globalEnv
}
