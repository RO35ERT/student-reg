// db/migrate.go

package db

import (
	"os"
	"student-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migrate() *gorm.DB {
    dsn := os.Getenv("DB_USER")+":"+ os.Getenv("DB_PASS") +"@tcp(localhost:3306)/StudentDB?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }

    db.AutoMigrate(&models.Student{})
    return db
}
