package main

func TestUpdate() {
	// update 只更新选择字段
	// updates 更新所有字段，此时有两种形式：一种为map，一种为结构体，结构体为零值，不参与更新
	//save 无论如何都更新，零值也更新
	var user []TestUser
	Db.Find(&user).Updates(TestUser{Name: "test", Age: 23})
}
