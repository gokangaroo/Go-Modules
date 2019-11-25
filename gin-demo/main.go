package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//全局中间件 允许跨域
func GlobalMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Next()
}

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "HelloWorld")
	})
	// /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		//获取firstname参数,firstname参数不存在，则使用默认的值 也就是c.DefaultQuery的第二参数值
		firstname := c.DefaultQuery("firstname", "Guest")
		//获取lastname参数，不设置默认值
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.POST("/form_post", func(c *gin.Context) {
		//获取post参数
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": message,
			"nick":    nick,
		})
	})
	router.Use(GlobalMiddleware) // listen and serve on 0.0.0.0:8080
	router.Run()
}
