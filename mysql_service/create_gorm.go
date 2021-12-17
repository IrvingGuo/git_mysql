package main

import "fmt"

func CreateTest() {
	dbres := Db.Create(&[]TestUser{
		{Name: "gsc", Age: 22},
		{Name: "qws", Age: 21},
		{Name: "erwe", Age: 20},
		{Name: "fret", Age: 19},
		{Name: "safg", Age: 18},
	})
	fmt.Println(dbres.Error, dbres.RowsAffected)
	if dbres.Error != nil {
		fmt.Println("创建失败")
	} else {
		fmt.Println("创建成功")
	}

}
