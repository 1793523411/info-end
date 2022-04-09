package table

import (
	"fmt"
	"time"

	allconst "info-end/const"
	"info-end/handler/table"
	"info-end/middleware"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RegistyUser(c *gin.Context) {
	body := table.User{}
	c.BindJSON(&body)
	data, err := table.InsertUser(allconst.Client, "user", body)
	fmt.Println(err, err != nil)
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
			"data": data.UserName,
		})
	}

}

func LoginUser(c *gin.Context) {
	body := table.User{}
	c.BindJSON(&body)
	res, err := table.Login(allconst.Client, body.UserName, body.PassWord)
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

func RefreshToken(c *gin.Context) {
	jwtCon := middleware.NewJWT()
	body := table.User{}
	c.BindJSON(&body)
	token, err := jwtCon.CreateToken(middleware.CustomClaims{
		UserName: body.UserName,
		Password: "refresh",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
	})
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
			"data": token,
		})
	}

}

func SaveUserInfo(c *gin.Context) {
	body := table.UserInfo{}
	c.BindJSON(&body)
	res, err := table.HandleUserInfo(allconst.Client, body)
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
func SearchUserInfo(c *gin.Context) {
	uid := c.Query("uid")
	username := c.Query("username")
	res, err := table.SearchUserInfo(allconst.Client, uid, username)
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