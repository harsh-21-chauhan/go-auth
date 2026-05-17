package models

import "gorm.io/gorm"

type User struct {
  gorm.Model              //  ID, CreatedAt,UpdatedAt,DeletedAt
  Email    string `json:"email" gorm:"unique;not null"`
  Password string `json:"-"`            // never send JSON
  Role     string `json:"role" gorm:"default:user"`
}

