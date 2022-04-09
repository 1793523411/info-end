package table

import (
	"fmt"
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
