package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "phpmyadmin:tzrsurya212@tcp(127.0.0.1:3306)/hehe-marketplace?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
