package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (o OrderController) GetList(c *gin.Context) {
	// cid := c.PostForm("cid")
	// name := c.DefaultPostForm("name", "wangwu")
	// param := make(map[string]interface{})
	// err := c.BindJSON(&param)
	Search := &Search{}
	err := c.BindJSON(&Search)
	if err == nil {
		// ReturnSuccess(c, 0, param["name"], param["cid"], 1)
		ReturnSuccess(c, 0, Search.Name, Search.Cid, 1)
		return
	}
	ReturnError(c, 4001, gin.H{"err": err})
	// ReturnSuccess(c, 0, cid, name, 1)
}
