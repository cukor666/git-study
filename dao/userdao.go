package dao

import "fmt"

type UserDao struct{}

func (u *UserDao) GetUser(id int) string {
	fmt.Println("根据ID获取用户信息")
	return "UserDao.GetUser"
}

func (u *UserDao) InsertUser(name string, age int) int {
	fmt.Println("插入用户信息")
	return 1
}
