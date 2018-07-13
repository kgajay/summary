package model

import "dao"

func Migrate() {
	testdb := dao.GetDb()
	testdb.AutoMigrate(&Product{}, &User{}, &Animal{})
}
