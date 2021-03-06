package main

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	UUID uint      `gorm: "primaryKey"`
	Time time.Time `gorm: "column: my_time"`
}
type TestUser struct {
	gorm.Model
	Name string `gorm:"default: qm"`
	Age  uint8  `gorm:"comment: 年龄"`
}

func TestUserCreate() {
	Db.AutoMigrate(&TestUser{})
}
