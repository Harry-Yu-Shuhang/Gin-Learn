package models

import "gin-learn/dao"

type User struct {
	ID       int
	Username string
}

func (User) TableName() string {
	return "user"
}

func GetUserSingle(id int) (User, error) {
	var user User
	err := dao.Db.Where("id =?", id).First(&user).Error //获取到的值给user，指针传递
	return user, err
}

func GetUserListTest(query_sentence string, query_variable any) ([]User, error) {
	var users []User
	// err := dao.Db.Where("id < ?", 3).Find(&users).Error //查询多个
	err := dao.Db.Where(query_sentence, query_variable).Find(&users).Error //查询多个
	return users, err
}

func AddUser(username string) (int, error) {
	user := User{Username: username}
	err := dao.Db.Create(&user).Error
	return user.ID, err
}

func UpdateUser(id int, username string) error {
	// 根据条件更新
	err := dao.Db.Model(&User{}).Where("id = ?", id).Update("username", username).Error
	// UPDATE users SET username=变量username, updated_at='2013-11-17 21:34:10' WHERE id=变量id;
	return err
}

func DeleteUser(id int) error {
	// 根据条件删除
	err := dao.Db.Delete(&User{}, "id = ?", id).Error
	// DELETE FROM users WHERE id = 变量id;
	return err
}
