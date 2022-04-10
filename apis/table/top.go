package table

import (
	"fmt"
	allconst "info-end/const"
	"info-end/handler/table"
	"info-end/middleware"
	"info-end/utils"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

//! 上传用户头像
func UploadTopImg(c *gin.Context) {
	rand.Seed(time.Now().Unix())
	file, _ := c.FormFile("file")
	claims, _ := c.Get("claims")
	username := claims.(*middleware.CustomClaims).UserName
	dirPeth := utils.GetCurrentAbPathByExecutable() + "/upload/"
	os.Mkdir(dirPeth, os.ModePerm)
	filename := username + "-" + fmt.Sprint(rand.Int63()) + "-" + file.Filename
	filepath := dirPeth + filename

	c.SaveUploadedFile(file, filepath)
	resPath, err := table.UploadTopImg(filename, filepath)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": "",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": resPath,
		})
	}
}

func CreateTopicRecord(c *gin.Context) {
	body := table.TopicRecord{}
	c.BindJSON(&body)
	res, err := table.CreateTopicRecord(allconst.Client, body)
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

func GetAllTopicRecord(c *gin.Context) {
	body := table.TopicSearchParams{}
	username, _ := c.Get("username")
	userType, _ := table.GetUserType(allconst.Client, username)
	c.BindJSON(&body)
	res, err := table.GetAllTopicRecord(allconst.Client, body, userType, username)
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

func SearchTopicRecordByRid(c *gin.Context) {
	rid := c.Query("rid")
	res, err := table.SearchTopicRecordByRid(allconst.Client, rid)
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

func UpdateTopicRecord(c *gin.Context) {
	body := table.TopicRecord{}
	c.BindJSON(&body)
	res, err := table.UpdateTopicRecord(allconst.Client, body)
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

func DelTopicRecord(c *gin.Context) {
	rid := c.Query("rid")
	username, _ := c.Get("username")
	res, err := table.DelTopicRecord(allconst.Client, rid, username)
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
