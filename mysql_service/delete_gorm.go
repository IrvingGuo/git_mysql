package main

func TestDel() {
	var user []TestUser
	Db.Where("name LIKE ?", "test").Delete(&user)
}
