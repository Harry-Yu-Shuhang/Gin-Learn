package main

import (
	router "gin-learn/router"
)

func main() {
	//封装路由
	r := router.Router()
	r.Run(":9999")
}
