package main

import (
	"database/sql"
	"time"
)

type Model struct {
	UUID uint      `gorm: "primaryKey"`
	Time time.Time `gorm: "column: my_time"`
}
type TestUser struct {
	Model        Model   `gorm:"embedded; embeddedPrefix: qq_"`
	Name         string  `gorm:"default: qm"`
	Email        *string `gorm:"not null"`
	Age          uint8   `gorm:"comment: 年龄"`
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
}

func TestUserCreate() {
	Db.AutoMigrate(&TestUser{})
}
