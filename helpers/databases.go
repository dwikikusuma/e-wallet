package helpers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"wallet/internal/model"
)

var DB *gorm.DB

func SetupMySQL() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", GetEnv("DB_USER", ""), GetEnv("DB_PASSWORD", ""), GetEnv("DB_HOST", "127.0.0.1"), GetEnv("DB_PORT", "3306"), GetEnv("DB_NAME", ""))
	logrus.Info("Connecting to db with dsn: ", dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connetct to databases: %e", err)

	}
	logrus.Info("Successfully connect to db")
	err = DB.AutoMigrate(&model.User{}, &model.UserSession{})

	if err != nil {
		log.Fatalf("failed to migrate databases: %e", err)
	} else {
		logrus.Info("Successfully migrate db")
	}
}
