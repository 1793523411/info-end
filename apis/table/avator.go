package table

import (
	allconst "info-end/const"
	"info-end/handler/table"

	"github.com/gin-gonic/gin"
)

func AvatorGet(c *gin.Context) {
	res, err := table.AvatorGet(allconst.Client)
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

func AvatorDel(c *gin.Context) {
	uid := c.Query("uid")
	err := table.AvatorDel(allconst.Client, uid)
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
			"data": "",
		})
	}
}
