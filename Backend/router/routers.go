package router

import (
	"gin-learn/config"
	"gin-learn/controllers"
	"gin-learn/pkg/logger"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine { //要大写，公有，别的地方要用
	r := gin.Default()

	//调用写日志函数,logger是自己写的
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	store, _ := redis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store)) //store指定存储方式，也就是redis
	// 1.	10：指定最大空闲连接数，即在连接池中 Redis 连接的最大空闲数量。
	// 2.	“tcp”：表示与 Redis 的连接方式是通过 TCP 协议。
	// 3.	config.RedisAddress：Redis 服务器的地址（包括 IP 和端口）。这个值通常从配置文件（config）中读取，例如 localhost:6379。
	// 4.	””：这个位置是 Redis 的密码字段（如果 Redis 没有设置密码，可以传空字符串）。
	// 5.	[]byte("secret")：加密密钥，用于对会话数据进行加密。这是一段字节数组（[]byte），防止会话数据被未授权访问。

	user := r.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.PostRegister)
		user.POST("/login", controllers.UserController{}.PostLogin) //用GET会把密码暴露在url中
	}

	player := r.Group("/player")
	{
		player.GET("/list/:aid", controllers.PlayerController{}.GetPlayersByAid)
	}

	vote := r.Group("/vote")
	{
		vote.POST("/add", controllers.VoteController{}.PostVote)
	}

	rank := r.Group("/rank")
	{
		rank.GET("/:aid", controllers.PlayerController{}.GetRankByAid)
	}

	return r
}
