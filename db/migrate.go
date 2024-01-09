package db

import (
	"student-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migrate() *gorm.DB {
    dsn := "root:tumbwerobert@tcp(localhost:3306)/StudentDB?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }

    db.AutoMigrate(&models.Student{})
    return db
}
