/*
gin框架实现统一鉴权与api版本控制
1.RESTful api 设计 （动作由请求方法决定）
2.路由分组实现api版本控制
3.中间件拦截请求实现统一鉴权
4.模型绑定与验证
参考视频：https://www.bilibili.com/video/BV1sw411m7YN P2
代码复现：CircleTAO
日期：2024年3月1日09:20:53
*/

package main

import (
	"modulename/src/gocode/golangproject2/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()
	//// 根据不同的请求方法形成不同的动作
	// r.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.POST("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "post pong",
	// 	})
	// })
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	routers.InitRouters(r)
	r.Run(":8080")
}
