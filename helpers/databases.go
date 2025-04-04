package bootstrap

import "fmt"

func SetupMySQL() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", GetEnv("DB_USER", ""), GetEnv("DB_PASSWORD", ""), GetEnv("DB_HOST", "127.0.0.1"), GetEnv("DB_PORT", "3306"), GetEnv("DB_NAME", ""))

}
