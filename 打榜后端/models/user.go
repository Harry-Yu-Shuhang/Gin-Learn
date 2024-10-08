package models

import (
	"gin-learn/dao"
	"time"
)

type User struct {
	ID         int    `json:"id"` //gorm默认处理方式，会把ID映射到mysql的id，即大写的变成小写。json是把结构体的数据映射进去，比如传入一个User结构体，里面全是大写的这种标签，json输出就是小写的这些标签
	Username   string `json:"username"`
	Password   string `json:"password"`
	AddTime    int64  `json:"add_time"`    //gorm默认的会把这种多个大写的映射到mysql中add_time,中间有一个下划线
	UpdateTime int64  `json:"update_time"` //也可以添加gorm参数指定映射的mysql列名，例如`gorm:"column:update_time"`，就把UpdateTime映射到mysql到update_time列了
}

func (User) TableName() string {
	return "user" //表名，Gorm中可以用TableName指定表名.必须用方法，不可以直接添加在结构体中
}

func GetUserInfoByUserName(username string) (User, error) {
	var user User
	err := dao.Db.Where("username = ?", username).First(&user).Error //获取到的值给user，指针传递
	// SELECT * FROM users WHERE username = 变量username ORDER BY id LIMIT 1;
	return user, err
}

func AddNewUser(username string, password string) (int, error) {
	user := User{Username: username, Password: password, AddTime: time.Now().Unix(), UpdateTime: time.Now().Unix()}
	err := dao.Db.Create(&user).Error
	//INSERT INTO user (username, password, add_time, update_time)
	//VALUES (变量username, 变量password, time.Now().Unix(), time.Now().Unix());
	//由于id是自增主键，因此会在后面加一个新的id，insert新的数据进去
	return user.ID, err
}

// func GetUserSingle(id int) (User, error) {
// 	var user User
// 	err := dao.Db.Where("id =?", id).First(&user).Error //获取到的值给user，指针传递
// 	return user, err
// }

// func GetUserListTest(query_sentence string, query_variable any) ([]User, error) {
// 	var users []User
// 	// err := dao.Db.Where("id < ?", 3).Find(&users).Error //查询多个
// 	err := dao.Db.Where(query_sentence, query_variable).Find(&users).Error //查询多个
// 	return users, err
// }

// func AddUser(username string) (int, error) {
// 	user := User{Username: username}
// 	err := dao.Db.Create(&user).Error
// 	return user.ID, err
// }

// func UpdateUser(id int, username string) error {
// 	// 根据条件更新
// 	err := dao.Db.Model(&User{}).Where("id = ?", id).Update("username", username).Error
// 	// UPDATE users SET username=变量username, updated_at='2013-11-17 21:34:10' WHERE id=变量id;
// 	return err
// }

// func DeleteUser(id int) error {
// 	// 根据条件删除
// 	err := dao.Db.Delete(&User{}, "id = ?", id).Error
// 	// DELETE FROM users WHERE id = 变量id;
// 	return err
// }
