package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginReq struct {
	UserName string `form:"user_name"`
	Pwd      string `form:"pwd"`
}

func Login(c *gin.Context) {
	req := loginReq{}
	c.Bind(&req) // 若模型不匹配，Bind()方法会直接将模型返回
	c.JSON(http.StatusOK, req)
	// c.BindQuery() // 如果是GET方式传参则用该方法操作
	// c.BindJSON()
	// err := c.ShouldBind() // 若模型不匹配，ShouldBind()方法会返回错误类型
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "login",
	// })
}
