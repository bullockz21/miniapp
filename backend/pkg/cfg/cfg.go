package cfg

import (
	"miniapp/pkg/logger"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Cfg struct {
	JwtSecretKey string `yaml:"jwt_secret_key"`
	Listen       struct {
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	Postgresql struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Database string `json:"database"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"Postgresql"`
}

var instance *Cfg
var once sync.Once

func GetConfig() *Cfg {
	once.Do(func() {
		logger := logger.GetLogger()
		logger.Infoln("read app configuration")
		instance = &Cfg{}
		err := cleanenv.ReadConfig("config.yml", instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Infoln(help)
			logger.Fatal(err)
		}
	})
	return instance
}
