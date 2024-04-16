package router

import (
	"golangproject4/controllers"
	"golangproject4/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 设置日志中间件
	r.Use(gin.LoggerWithConfig((logger.LoggerToFile())))
	r.Use(logger.Recover)

	user := r.Group("/user")
	{
		user.GET("/info/:id", controllers.UserController{}.GetUserInfo)
		user.POST("/list", controllers.UserController{}.GetList)
		user.POST("/add", controllers.UserController{}.AddUser)
		// user.PUT("/add", func(c *gin.Context) {
		// 	c.String(http.StatusOK, "user add")
		// })
		user.POST("/update", controllers.UserController{}.UpdateUser)
		user.POST("/delete", controllers.UserController{}.DeleteUser)
		user.GET("/list/test", controllers.UserController{}.GetUserList)

		user.DELETE("/delete", func(c *gin.Context) {
			c.String(http.StatusOK, "user delete")
		})
	}

	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}
	return r
}
