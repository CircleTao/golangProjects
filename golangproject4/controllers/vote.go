package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"golangproject4/cache"
	"golangproject4/models"
)

type VoteController struct{}

func (v VoteController) AddVote(c *gin.Context) {
	userIdStr := c.DefaultPostForm("userId", "0")
	playerIdStr := c.DefaultPostForm("playerId", "0")
	userId, _ := strconv.Atoi(userIdStr)
	playerId, _ := strconv.Atoi(playerIdStr)
	if userId == 0 || playerId == 0 {
		ReturnError(c, 4001, "please input correct userId and playerId")
		return
	}
	user, _ := models.GetUserInfoById(userId)
	if user.Id == 0 {
		ReturnError(c, 4001, "userId not exist!")
		return
	}
	player, _ := models.GetPlayerInfoById(playerId)
	if player.Id == 0 {
		ReturnError(c, 4001, "playerId not exist!")
		return
	}
	vote, _ := models.GetVoteInfoById(userId, playerId)
	if vote.Id != 0 {
		ReturnError(c, 4001, "you have voted!")
		return
	}

	rs, err := models.AddVote(userId, playerId)
	if err == nil {
		models.UpdatePlayerScore(playerId)
		redisKey := "ranking:" + strconv.Itoa(player.Aid)
		cache.Rdb.ZIncrBy(cache.Rtcx, redisKey, 1, strconv.Itoa(playerId))
		ReturnSuccess(c, 0, "vote success!", rs, 1)
		return
	}
	ReturnError(c, 4004, "please contact management!")
}
