package cfg

import (
	"log"
	"miniapp/pkg/logger"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Cfg struct {
	JwtSecretKey string `env:"JWT_SECRET_KEY"`
	Listen       struct {
		BindIP string `env:"LISTEN_BIND_IP" env-default:"127.0.0.1"`
		Port   string `env:"LISTEN_PORT" env-default:"8080"`
	}
	Postgresql struct {
		Host     string `env:"POSTGRESQL_HOST"`
		Port     string `env:"POSTGRESQL_PORT"`
		Database string `env:"POSTGRESQL_DATABASE"`
		Username string `env:"POSTGRESQL_USERNAME"`
		Password string `env:"POSTGRESQL_PASSWORD"`
	}
}

var instance *Cfg
var once sync.Once

// func GetConfig() *Cfg {
// 	once.Do(func() {
// 		logger := logger.GetLogger()
// 		logger.Infoln("read app configuration")
// 		instance = &Cfg{}
// 		err := cleanenv.ReadConfig("config.yml", instance)
// 		if err != nil {
// 			help, _ := cleanenv.GetDescription(instance, nil)
// 			logger.Infoln(help)
// 			logger.Fatal(err)
// 		}
// 	})
// 	return instance
// }

func GetConfigEnv() *Cfg {
	once.Do(func() {
		logger := logger.GetLogger()
		logger.Infoln("read app configuration")
		instance = &Cfg{}
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("error open .env file")
		}
		err = cleanenv.ReadConfig(".env", instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Infoln(help)
			logger.Fatal(err)
		}
	})
	return instance
}
