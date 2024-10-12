package models

import (
	"gin-learn/dao"

	"gorm.io/gorm"
)

type Player struct {
	ID          int    `json:"id"`
	Aid         int    `json:"aid"`
	Ref         string `json:"ref"`
	Nickname    string `json:"nick_name"`
	Declaration string `json:"declaration"`
	Avatar      string `json:"avatar"`
	Score       int    `json:"score"`
}

func (Player) TableName() string {
	return "player"
}

func GetPlayersByAid(aid int, sort string) ([]Player, error) {
	var players []Player
	err := dao.Db.Where("aid = ?", aid).Order(sort).Find(&players).Error
	return players, err
}

func GetPlayerInfoByID(id int) (Player, error) {
	var player Player
	err := dao.Db.Where("id = ?", id).First(&player).Error
	return player, err
}

func IncreasePlayerScore(id int) {
	var player Player
	err := dao.Db.Model(&player).Where("id = ?", id).UpdateColumn("score", gorm.Expr("score + ?", 1)).Error
	// UPDATE "player" SET "score" = score + 1 WHERE "id" = id;
	if err != nil {
		panic(err)
	}
}
