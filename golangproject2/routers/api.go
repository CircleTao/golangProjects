package routers

import (
	"modulename/src/gocode/golangproject2/web"

	"github.com/gin-gonic/gin"
)

func initApi(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/ping", web.Ping) // 传入方法web.ping而非调用方法，因此无须()
	v1.POST("/login", web.Login)
	v1.POST("/register", web.Register)
}
