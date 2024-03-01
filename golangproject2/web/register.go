package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type registerReq struct {
	UserName string `form:"user_name" binding:"required"`
	Pwd      string `form:"pwd" binding:"required"`
	Phone    string `form:"phone" binding:"required,e164"`
	Email    string `form:"email" binding:"omitempty,email"`
}

func Register(c *gin.Context) {
	req := &registerReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, req)
	c.JSON(http.StatusOK, gin.H{
		"message": "register",
	})
}
