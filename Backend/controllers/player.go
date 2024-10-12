package controllers

import (
	"gin-learn/cache"
	"gin-learn/config"
	"gin-learn/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PlayerController struct{}

func (p PlayerController) GetPlayersByAid(ctx *gin.Context) {
	aidStr := ctx.Param("aid")
	aid, _ := strconv.Atoi(aidStr)
	err_id := config.GetPlayersErr

	rs, err := models.GetPlayersByAid(aid, "id asc") //排序方式为id升序
	if err != nil {
		ReturnError(ctx, err_id, "该aid不存在")
		return
	}
	ReturnSuccess(ctx, 0, "获取成功", rs, 1)
}

func (p PlayerController) GetRankByAid(ctx *gin.Context) {
	//测试代码，看看能不能加到redis里面
	// err := cache.Rdb.Set(cache.Rctx, "name", "张三", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	aidStr := ctx.Param("aid")
	aid, _ := strconv.Atoi(aidStr)
	err_id := config.GetRankingErr

	redisKey := "rank:" + aidStr                                         //redis的key有层级的话，用:隔开就可以
	rs, err := cache.Rdb.ZRevRange(cache.Rctx, redisKey, 0, -1).Result() //从0开始到-1也就是最后一个元素，倒序排列。ZRevRange返回的事Cmd对象，常用Result获取结果。
	//rs存的是member id
	if err == nil && len(rs) > 0 {
		var players []models.Player
		for _, value := range rs {
			id, _ := strconv.Atoi(value)
			rsInfo, _ := models.GetPlayerInfoByID(id) //后续可以优化，把参赛选手详情也存到redis里面，减少mysql查询
			if rsInfo.ID > 0 {
				players = append(players, rsInfo)
			}
		}
		ReturnSuccess(ctx, 0, "获取成功", players, 1)
		//如果redis获取到了，就直接返回。没获取到继续去mysql获取,再返回redis中
		return
	}

	rsSql, err := models.GetPlayersByAid(aid, "score desc") //排序方式为score降序
	if err != nil {
		ReturnError(ctx, err_id, "该aid不存在")
		return
	}
	for _, value := range rsSql {
		cache.Rdb.ZAdd(cache.Rctx, redisKey, cache.Zscore(value.ID, value.Score)).Err() //把数据存入redis
	}
	//设置过期时间
	cache.Rdb.Expire(cache.Rctx, redisKey, 24*time.Hour) //设置24小时过期
	ReturnSuccess(ctx, 0, "获取成功", rsSql, 1)
}
