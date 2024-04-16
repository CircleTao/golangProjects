package controllers

import (
	"golangproject4/models"
	"golangproject4/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	idStr := c.Param("id")
	name := c.Param("name")

	id, _ := strconv.Atoi(idStr)
	user, _ := models.GetUserTest(id)
	ReturnSuccess(c, 0, name, user, 1)
}

func (u UserController) GetList(c *gin.Context) {
	num1 := 1
	num2 := 0
	num3 := num1 / num2
	logger.Write("log info", "user")
	// ReturnError(c, 4004, "get list error!")
	ReturnError(c, 4004, num3)
}

func (u UserController) AddUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	id, err := models.AddUser(username)
	if err != nil {
		ReturnError(c, 4002, "add user error!")
		return
	}
	ReturnSuccess(c, 0, "add user success!", id, 1)
}

func (u UserController) UpdateUser(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "")
	username := c.DefaultPostForm("username", "")
	id, _ := strconv.Atoi(idStr)
	models.UpdateUser(id, username)
	ReturnSuccess(c, 0, "update user success!", id, 1)
}

func (u UserController) DeleteUser(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	err := models.DeleteUser(id)
	if err != nil {
		ReturnError(c, 4003, "delete user error!")
		return
	}
	ReturnSuccess(c, 0, "delete user success!", id, 1)
}

func (u UserController) GetUserList(c *gin.Context) {
	users, err := models.GetUserList()
	if err != nil {
		ReturnError(c, 4004, "get user list error!")
		return
	}
	ReturnSuccess(c, 0, "get user list success!", users, 1)
}
