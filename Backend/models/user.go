package models

import "gin-learn/dao"

type User struct {
	ID       int
	Username string
}

func (User) TableName() string {
	return "user"
}

func GetUserTest(id int) (User, error) {
	var user User
	err := dao.Db.Where("id =?", id).First(&user).Error //获取到的值给user，指针传递
	return user, err
}
