package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./assets")
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"status": "ok",
		})
	})

	// path parameter
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		fullPath := c.FullPath()
		// response content-type: text/plain; charset=utf-8
		c.String(http.StatusOK, "Hello %s %s", name, fullPath)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	// : *都是参数匹配，有什么区别？*会包含前面的/。整体认为iris的写法更友好
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		fullPath := c.FullPath()
		c.String(http.StatusOK, "%v %v", message, fullPath)
	})

	// by default, it serves on :8080 unless a PORT environment variable was defined.
	r.Run()
}

func RegisterRouting(r *gin.RouterGroup) {
	v1 := r.Group("/oauth2/")
	{
		v1.POST("/token")
	}

}
