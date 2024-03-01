package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {

	// 初始化API路由（不需要鉴权）
	initApi(r)

	// 初始化课程路由
	initCourse(r)
}
