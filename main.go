package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yaolei313/gin-study/internal/note"
	"io"
	"net/http"
	"os"
)

func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f) 只写文件
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// gin.Default With the Logger and Recovery middleware already attached
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		// type assertion https://go.dev/tour/methods/15
		errStr, ok := err.(string)
		if ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", errStr))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))


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
	// : *都是参数匹配，有什么区别？*会包含前面的/字符。
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		fullPath := c.FullPath()
		c.String(http.StatusOK, "%v %v", message, fullPath)
	})

	r.GET("/panic", func(c *gin.Context) {
		panic("test panic")
	})

	registerDraftRouting(r)


	// by default, it serves on :8080 unless a PORT environment variable was defined.
	r.Run()
}

func RegisterLoginRouting(r *gin.Engine) {
	v1 := r.Group("/user/passport/")
	{
		v1.POST("/login_by_password", )
	}
}

func registerDraftRouting(r *gin.Engine) {
	draftGroup := r.Group("/note/draft")

	draftGroup.GET("/:id", note.GetDraft)
	draftGroup.POST("/save", note.SaveDraft)
	draftGroup.POST("/del", note.DeleteDraft)
}
