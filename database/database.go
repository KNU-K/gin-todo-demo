package database

import (
	"log"
	"todo-app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Connect() {
    var err error
    dsn := "root:root@tcp(127.0.0.1:3306)/go_todo?charset=utf8mb4&parseTime=True&loc=Local"
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("DB 연결 실패:", err)
    } else {
        log.Println("DB 연결 성공")
    }

    // Todo 모델 마이그레이션
    err = DB.AutoMigrate(&models.Todo{})
    if err != nil {
        log.Fatal("DB 마이그레이션 실패:", err)
    } else {
        log.Println("DB 마이그레이션 성공")
    }
}