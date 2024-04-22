package controllers

import (
	"strconv"
	"time"

	"golangproject4/cache"
	"golangproject4/models"

	"github.com/gin-gonic/gin"
)

type PlayerController struct{}

func (p PlayerController) GetPlayerInfo(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)
	// player, err := models.GetPlayerInfoByAid(aid)
	player, err := models.GetPlayersSort(aid, "id desc")
	if err != nil {
		ReturnError(c, 4004, "get player info error!")
		return
	}
	ReturnSuccess(c, 0, "success", player, 1)
}

func (p PlayerController) GetRanking(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)

	redisKey := "ranking:" + aidStr
	rs, err := cache.Rdb.ZRevRange(cache.Rtcx, redisKey, 0, -1).Result()
	if err == nil && len(rs) > 0 {
		var players []models.Player
		for _, value := range rs {
			id, _ := strconv.Atoi(value)
			rsInfo, _ := models.GetPlayerInfoById(id)
			if rsInfo.Id > 0 {
				players = append(players, rsInfo)
			}
		}
		ReturnSuccess(c, 0, "success", players, 1)
		return
	}

	player, errDb := models.GetPlayersSort(aid, "score desc")
	if errDb == nil {
		for _, value := range player {
			cache.Rdb.ZAdd(cache.Rtcx, redisKey, cache.Zscore(value.Id, value.Score))
		}
		cache.Rdb.Expire(cache.Rtcx, redisKey, 24*time.Hour)
		ReturnSuccess(c, 0, "success", player, 1)
		return
	}
	ReturnError(c, 4004, "get player info error!")
}
