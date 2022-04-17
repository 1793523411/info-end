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

type UserLogin struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

// LoginUser 用户登录
// @Summary 用户登录
// @Description 用户登录
// @Tags user
// @Produce application/json
// @Param  body body UserLogin true "用户登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} table.loginInfo
// @Router /user/login [post]
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

type SwggerRefeshToken struct {
	UserName string `json:"username"`
}

type RefreshTokenRes struct {
	Token string `json:"token"`
}

// SearchUserInfo 刷新用户token
// @Summary 刷新用户token!
// @Description 刷新用户token!!
// @Tags user
// @Produce application/json
// @Param object query SwggerRefeshToken true "查询参数"
// @Success 200 {object} RefreshTokenRes
// @Router /user/refresh_token [get]
func RefreshToken(c *gin.Context) {
	jwtCon := middleware.NewJWT()
	username := c.Query("username")
	token, err := jwtCon.CreateToken(middleware.CustomClaims{
		UserName: username,
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
			"data": RefreshTokenRes{
				Token: token,
			},
		})
	}

}

// SaveUserInfo 更新用户信息
// @Summary 更新用户信息
// @Description 更新用户信息
// @Tags user
// @Produce application/json
// @Param  body body table.UserInfo true "用户登录参数"
// @Param token header string true "Bearer 用户令牌"
// @Success 200 {object} table.UserInfo
// @Router /api/v1/save_user_info [post]
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

type SwggerSearchUserInfo struct {
	Uid string `json:"uid"`
}

// SearchUserInfo 查询用户信息
// @Summary 查询用户信息!
// @Description 查询用户信息!!
// @Tags user
// @Produce application/json
// @Param object query SwggerSearchUserInfo true "查询参数"
// @Param token header string true "Bearer 用户令牌"
// @Success 200 {object} table.UserInfo
// @Router /api/v1/search_user_info [get]
func SearchUserInfo(c *gin.Context) {
	uid := c.Query("uid")
	username, _ := c.Get("username")
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
