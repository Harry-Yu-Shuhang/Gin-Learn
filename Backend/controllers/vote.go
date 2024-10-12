package controllers

import (
	"gin-learn/cache"
	"gin-learn/config"
	"gin-learn/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VoteController struct{}

func (v VoteController) PostVote(ctx *gin.Context) {
	//获取用户id 选手id
	userIdStr := ctx.DefaultPostForm("user_id", "0")
	playerIdStr := ctx.DefaultPostForm("player_id", "0")
	userId, _ := strconv.Atoi(userIdStr)
	playerId, _ := strconv.Atoi(playerIdStr)
	err_id := config.IncreaseVoteErr

	if userId == 0 {
		ReturnError(ctx, err_id, "用户id错误,请联系管理员")
		return
	} else if playerId == 0 {
		ReturnError(ctx, err_id, "选手id错误,请联系管理员")
		return
	}
	user, _ := models.GetUserInfoByID(userId)
	if user.ID == 0 {
		ReturnError(ctx, err_id, "投票用户不存在")
		return
	}
	player, _ := models.GetPlayerInfoByID(playerId)
	if player.ID == 0 {
		ReturnError(ctx, err_id, "参赛选手不存在")
		return
	}
	vote, _ := models.GetVoteInfo(userId, playerId)
	if vote.ID != 0 {
		ReturnError(ctx, err_id, "您已经给这位选手投过票了")
		return
	}

	rs, err := models.PostVote(userId, playerId)
	if err != nil {
		ReturnError(ctx, err_id, "投票失败,请联系管理员")
		return
	}
	//增加投票数
	models.IncreasePlayerScore(playerId)
	//更新数据库以后同时也要更新redis，否则刷新出来还是redis的值，排行榜没有变化
	redisKey := "rank:" + strconv.Itoa(player.Aid)
	cache.Rdb.ZIncrBy(cache.Rctx, redisKey, 1, strconv.Itoa(playerId))
	//这行代码在 Redis 中的有序集合（redisKey）中，将 playerId 转换为字符串后作为成员，并将该成员的分数增加 1。如果该成员之前不存在，会将其添加到集合中，初始分数为 1。
	ReturnSuccess(ctx, 0, "投票成功", rs, 1)
}
