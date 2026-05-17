package config

import (
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  "github.com/harsh-21-chauhan/go-auth/models"
)

var DB *gorm.DB  

func ConnectDB() {

  db, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})


  
  if err != nil {
    panic("Failed to connect to database: " + err.Error())
  }

 
  db.AutoMigrate(&models.User{})

  DB = db
}