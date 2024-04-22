package controllers

import (
	"golangproject4/models"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

type UserApi struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (u UserController) Register(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmpassword := c.DefaultPostForm("confirmPassword", "")
	if username == "" || password == "" || confirmpassword == "" {
		ReturnError(c, 4001, "please input correct info!")
		return
	}

	if password != confirmpassword {
		ReturnError(c, 4001, "password not match!")
		return
	}
	user, err := models.GetUserInfoByUserName(username)
	if user.Id != 0 {
		ReturnError(c, 4001, "username exist!")
		return
	}

	_, err = models.AddUser(username, EncryMd5(password))
	if err != nil {
		ReturnError(c, 4001, "register error!")
		return
	}
	ReturnSuccess(c, 0, "register success!", nil, 1)
}

func (u UserController) Login(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	if username == "" || password == "" {
		ReturnError(c, 4001, "please input correct info!")
		return
	}

	user, _ := models.GetUserInfoByUserName(username)
	if user.Id == 0 {
		ReturnError(c, 4001, "username or password not correct!")
		return
	}

	if user.Password != EncryMd5(password) {
		ReturnError(c, 4001, "username or password not correct!")
		return
	}

	session := sessions.Default(c)
	session.Set("login:"+strconv.Itoa(user.Id), user.Id)
	session.Save()
	err := session.Save()
	if err != nil {
		ReturnError(c, 500, "failed to save session")
		return
	}
	data := UserApi{Id: user.Id, Username: user.Username}
	ReturnSuccess(c, 0, "login success!", data, 1)
}

// func (u UserController) GetUserInfo(c *gin.Context) {
// 	idStr := c.Param("id")
// 	name := c.Param("name")

// 	id, _ := strconv.Atoi(idStr)
// 	user, _ := models.GetUserTest(id)
// 	ReturnSuccess(c, 0, name, user, 1)
// }

// func (u UserController) GetList(c *gin.Context) {
// 	num1 := 1
// 	num2 := 0
// 	num3 := num1 / num2
// 	logger.Write("log info", "user")
// 	// ReturnError(c, 4004, "get list error!")
// 	ReturnError(c, 4004, num3)
// }

// func (u UserController) AddUser(c *gin.Context) {
// 	username := c.DefaultPostForm("username", "")
// 	id, err := models.AddUser(username)
// 	if err != nil {
// 		ReturnError(c, 4002, "add user error!")
// 		return
// 	}
// 	ReturnSuccess(c, 0, "add user success!", id, 1)
// }

// func (u UserController) UpdateUser(c *gin.Context) {
// 	idStr := c.DefaultPostForm("id", "")
// 	username := c.DefaultPostForm("username", "")
// 	id, _ := strconv.Atoi(idStr)
// 	models.UpdateUser(id, username)
// 	ReturnSuccess(c, 0, "update user success!", id, 1)
// }

// func (u UserController) DeleteUser(c *gin.Context) {
// 	idStr := c.DefaultPostForm("id", "")
// 	id, _ := strconv.Atoi(idStr)
// 	err := models.DeleteUser(id)
// 	if err != nil {
// 		ReturnError(c, 4003, "delete user error!")
// 		return
// 	}
// 	ReturnSuccess(c, 0, "delete user success!", id, 1)
// }

// func (u UserController) GetUserList(c *gin.Context) {
// 	users, err := models.GetUserList()
// 	if err != nil {
// 		ReturnError(c, 4004, "get user list error!")
// 		return
// 	}
// 	ReturnSuccess(c, 0, "get user list success!", users, 1)
// }
