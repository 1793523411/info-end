package router

import (
	"fmt"

	"info-end/apis/other"
	"info-end/apis/table"
	_ "info-end/docs"
	"info-end/middleware"

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("这是中间件")
	}
}

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.GET("/", middleware.JWTAuth(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello!",
			"code":    0,
		})
	})
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	r.GET("/test", other.SearchDemo)

	user := r.Group("/user")
	{
		user.POST("/registy", table.RegistyUser)
		user.POST("/login", table.LoginUser)
		user.GET("/refresh_token", table.RefreshToken)
	}

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWTAuth())
	{
		apiv1.POST("/save_user_info", table.SaveUserInfo)
		apiv1.GET("/search_user_info", table.SearchUserInfo)
		apiv1.POST("/upload_user_avator", table.UploadUserAvator)

		apiv1.POST("/avator_add", table.AvatorAdd)
		apiv1.GET("/avator_get", table.AvatorGet)
		apiv1.GET("/avator_del", table.AvatorDel)

		apiv1.POST("/upload_video", table.VideoUpload)
		apiv1.POST("/upload_video_img", table.VideoImgUpload)
		apiv1.POST("/create_video_record", table.CreateVideoRecord)
		apiv1.POST("/get_video_list", table.GetAllVideoRecord)
		apiv1.GET("/search_video_list", table.SearchVideorecordByRid)
		apiv1.POST("/update_video_list", table.UpdateVideoRecord)
		apiv1.GET("/delete_video_record", table.DelVideoRecord)

		apiv1.POST("/upload_top_img", table.UploadTopImg)
		apiv1.POST("/create_topic_record", table.CreateTopicRecord)
		apiv1.POST("/get_topic_list", table.GetAllTopicRecord)
		apiv1.GET("/search_topic_list", table.SearchTopicRecordByRid)
		apiv1.POST("/update_topic_list", table.UpdateTopicRecord)
		apiv1.GET("/delete_topic_record", table.DelTopicRecord)

		apiv1.POST("/getList", func(c *gin.Context) {
			info, exit := c.Get("claims")
			if exit {
				c.JSON(200, gin.H{
					"message": info,
					"code":    0,
				})
			} else {
				c.JSON(500, gin.H{
					"message": info,
					"code":    1,
				})
			}
		})
	}
	r.Run(":8080")
}
