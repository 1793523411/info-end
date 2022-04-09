package table

import (
	"fmt"
	allconst "info-end/const"
	"info-end/middleware"
	"info-end/utils"
	"math/rand"
	"os"
	"time"

	"info-end/handler/table"

	"github.com/gin-gonic/gin"
)

//! 上传用户头像
func UploadUserAvator(c *gin.Context) {
	rand.Seed(time.Now().Unix())
	file, _ := c.FormFile("file")
	claims, _ := c.Get("claims")
	username := claims.(*middleware.CustomClaims).UserName
	dirPeth := utils.GetCurrentAbPathByExecutable() + "/upload/"
	os.Mkdir(dirPeth, os.ModePerm)
	filename := username + "-" + fmt.Sprint(rand.Int63()) + "-" + file.Filename
	filepath := dirPeth + filename

	c.SaveUploadedFile(file, filepath)
	resPath, err := table.UploadUserAvator(filename, filepath)
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

func AvatorAdd(c *gin.Context) {
	rand.Seed(time.Now().Unix())
	file, _ := c.FormFile("file")
	dirPeth := utils.GetCurrentAbPathByExecutable() + "/upload/"
	os.Mkdir(dirPeth, os.ModePerm)
	filename := "default-" + fmt.Sprint(rand.Int63()) + "-" + file.Filename
	filepath := dirPeth + filename

	c.SaveUploadedFile(file, filepath)
	resPath, _ := table.UploadUserAvator(filename, filepath)
	res, err := table.AvatorAdd(allconst.Client, resPath)
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
			"data": res,
		})
	}
}

func VideoUpload(c *gin.Context) {
	rand.Seed(time.Now().Unix())
	file, _ := c.FormFile("file")
	dirPeth := utils.GetCurrentAbPathByExecutable() + "/upload/"
	os.Mkdir(dirPeth, os.ModePerm)
	filename := "video-" + fmt.Sprint(rand.Int63()) + "-" + file.Filename
	filepath := dirPeth + filename

	c.SaveUploadedFile(file, filepath)
	resPath, _ := table.UploadUserVideo(filename, filepath)
	res, err := table.VideoUpload(allconst.Client, resPath)
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
			"data": res,
		})
	}
}

func VideoImgUpload(c *gin.Context) {
	rand.Seed(time.Now().Unix())
	file, _ := c.FormFile("file")
	dirPeth := utils.GetCurrentAbPathByExecutable() + "/upload/"
	os.Mkdir(dirPeth, os.ModePerm)
	filename := "video-img-" + fmt.Sprint(rand.Int63()) + "-" + file.Filename
	filepath := dirPeth + filename

	c.SaveUploadedFile(file, filepath)
	resPath, _ := table.UploadUserVideo(filename, filepath)
	res, err := table.VideoImgUpload(allconst.Client, resPath)
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
			"data": res,
		})
	}
}
