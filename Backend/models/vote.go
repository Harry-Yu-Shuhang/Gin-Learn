package models

import (
	"gin-learn/dao"
	"time"
)

type Vote struct {
	ID       int   `json:"id"` //gorm默认处理方式，会把ID映射到mysql的id，即大写的变成小写。json是把结构体的数据映射进去，比如传入一个User结构体，里面全是大写的这种标签，json输出就是小写的这些标签
	UserId   int   `json:"user_id"`
	PlayerId int   `json:"player_id"`
	AddTime  int64 `json:"add_time"` //gorm默认的会把这种多个大写的映射到mysql中add_time,中间有一个下划线
}

func (Vote) TableName() string {
	return "vote"
}

func GetVoteInfo(userId int, playerId int) (Vote, error) {
	var vote Vote
	err := dao.Db.Where("user_id =? AND player_id =?", userId, playerId).First(&vote).Error
	return vote, err
}

func PostVote(userId int, playerId int) (int, error) {
	vote := Vote{UserId: userId, PlayerId: playerId, AddTime: time.Now().Unix()}
	err := dao.Db.Create(&vote).Error
	return vote.ID, err
}
