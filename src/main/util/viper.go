package util

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type DBStruct struct {
	Host   string `mapstructure:"DB_HOST"`
	Port   string `mapstructure:"DB_PORT"`
	User   string `mapstructure:"DB_USER"`
	Pass   string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`
}

func LoadENV(path string) (db DBStruct, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	// read config
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error read .env file")
		log.Fatal(err.Error())
		os.Exit(1)
	}

	// unmarshal config
	err = viper.Unmarshal(&db)
	if err != nil {
		log.Fatal("Error unmarshal .env file")
		log.Fatal(err.Error())
		os.Exit(1)
	}

	return 
}
