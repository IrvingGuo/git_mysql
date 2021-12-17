package main

import "fmt"

type UserInfo struct {
	Name string
	Age  uint8
}

func TestFind() {
	var UserInfo []UserInfo
	Db.Model(&TestUser{}).Find(&UserInfo)
	fmt.Println(UserInfo)
}
