package util

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	env, err := LoadENV("../..")
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("host=%s user=%s pass=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", env.Host, env.User, env.Pass, env.DBName, env.Port)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func AuthMigrateDB(db *gorm.DB) error {
	return db.AutoMigrate()
}
