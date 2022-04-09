package table

import (
	allconst "info-end/const"
	"info-end/handler/table"

	"github.com/gin-gonic/gin"
)

func CreateVideoRecord(c *gin.Context) {
	body := table.VideoRecord{}
	c.BindJSON(&body)
	res, err := table.CreateVideoRecord(allconst.Client, body)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": res,
		})
	}
}

func GetAllVideoRecord(c *gin.Context) {
	body := table.SearchParams{}
	username, _ := c.Get("username")
	userType, _ := table.GetUserType(allconst.Client, username)
	c.BindJSON(&body)
	res, err := table.GetAllVideoRecord(allconst.Client, body, userType, username)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": res,
		})
	}
}

func SearchVideorecordByRid(c *gin.Context) {
	rid := c.Query("rid")
	res, err := table.SearchVideorecordByRid(allconst.Client, rid)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": res,
		})
	}
}

func UpdateVideoRecord(c *gin.Context) {
	body := table.VideoRecord{}
	c.BindJSON(&body)
	res, err := table.UpdateVideoRecord(allconst.Client, body)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": res,
		})
	}
}

func DelVideoRecord(c *gin.Context) {
	rid := c.Query("rid")
	username, _ := c.Get("username")
	res, err := table.DelVideoRecord(allconst.Client, rid, username)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": res,
		})
	}
}
