package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	Id        int    `gorm:"primaryKey" json:"id"`
	NamaUser  string `gorm:"type:varchar(255)" json:"nama_user"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
}

func ConnectDB() {
	dsn := "host=localhost user=postgres password=root dbname=go-learn port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	DB = db
}
