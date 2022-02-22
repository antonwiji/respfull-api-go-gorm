package main

import (
	"albums/gin/api"
	"albums/gin/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/api_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Person{})

	person := entity.Person{}
	person.Kelas = "XII"
	person.Nama = "Anton Wijaya"

	err = db.Create(&person).Error

	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/", api.ShowApi)
	router.GET("/siswa/:nama", api.ShowParams)
	router.GET("/query", api.QueryApi)

	router.Run()

}
