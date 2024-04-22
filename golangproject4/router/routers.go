package router

import (
	"golangproject4/config"
	"golangproject4/controllers"
	"golangproject4/pkg/logger"

	"github.com/gin-contrib/sessions"
	sessions_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 设置日志中间件
	r.Use(gin.LoggerWithConfig((logger.LoggerToFile())))
	r.Use(logger.Recover)
	store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	user := r.Group("/user")
	{
		// user.GET("/info/:id", controllers.UserController{}.GetUserInfo)
		// user.POST("/list", controllers.UserController{}.GetList)
		// user.POST("/add", controllers.UserController{}.AddUser)
		// // user.PUT("/add", func(c *gin.Context) {
		// // 	c.String(http.StatusOK, "user add")
		// // })
		// user.POST("/update", controllers.UserController{}.UpdateUser)
		// user.POST("/delete", controllers.UserController{}.DeleteUser)
		// user.GET("/list/test", controllers.UserController{}.GetUserList)

		// user.DELETE("/delete", func(c *gin.Context) {
		// 	c.String(http.StatusOK, "user delete")
		// })
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("/login", controllers.UserController{}.Login)
	}

	player := r.Group("/player")
	{
		player.POST("/info", controllers.PlayerController{}.GetPlayerInfo)
	}

	vote := r.Group("/vote")
	{
		vote.POST("/add", controllers.VoteController{}.AddVote)
	}

	r.POST("/ranking", controllers.PlayerController{}.GetRanking)
	// order := r.Group("/order")
	// {
	// 	order.POST("/list", controllers.OrderController{}.GetList)
	// }
	return r
}
