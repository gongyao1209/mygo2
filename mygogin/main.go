package main

import (
	"mygo2/db"
	_"mygo2/db"

	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main()  {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":"pong",
		})
	})

	v1 := r.Group("v1")
	paramQueryMap2(v1)

	paramInPath(r)

	paramInPath2(r)

	paramQuery1(r)

	//POST请求
	paramFromPost(r)
	//Query和Post联合
	paramQueryPost(r)
	//使用map传递参数
	paramQueryMap(r)

	// 上传文件
	r.POST("/upload", uploadFile)

	// 上传文件
	r.POST("/upload1", uploadFiles)

	r.Run(":12345")
}

/*

curl -X POST \
  http://127.0.0.1:12345/upload \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data' \
  -H 'postman-token: 8455203a-c358-16e1-e005-e8d881429acc' \
  -F file=@go.jpg

 */
func uploadFile(ctx *gin.Context)  {
	f, _ := ctx.FormFile("file")

	fmt.Println(f.Filename)
	//

	ctx.JSON(http.StatusOK, gin.H{"file_name":f.Filename})
}


/**
	curl -X POST \
  http://127.0.0.1:12345/upload1 \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data' \
  -H 'postman-token: 8bd3b5eb-fde4-c3ec-f993-fcf48ee5243b' \
  -F 'upload[]=@go.jpg' \
  -F 'upload[]=@QQ20180521-0.JPG'
 */
func uploadFiles(ctx *gin.Context)  {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status":"error"})
		return
	}

	files := form.File["upload[]"]

	for _, f := range files {

		fmt.Println(f.Filename)
	}

	ctx.JSON(http.StatusOK, gin.H{"status":"success"})
}

//参数在path里面
func paramInPath(router *gin.Engine)  {
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")

		strconv.Atoi(name)

		//TODO Do Something

		c.String(http.StatusOK, "Hello %s", name)
	})
}

//这个路由将指向 /user/john/ 和 /user/john/send
// 不会匹配 /user/john
func paramInPath2(router *gin.Engine)  {
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		c.String(http.StatusOK, "Hello %s, %s", name, action)
	})
}

//使用query传递参数
func paramQuery1(router *gin.Engine)  {
	router.GET("welcome", func(c *gin.Context) {
		//默认query值
		f_name := c.DefaultQuery("firstname", "gongyao")
		//获取lastname
		l_name := c.Query("lastname")

		c.String(http.StatusOK, "firstname: %s, lastname: %s", f_name, l_name)
	})
}


//使用post
func paramFromPost(router *gin.Engine)  {
	router.POST("from_post", func(c *gin.Context) {
		name := c.PostForm("name")
		age := c.DefaultPostForm("age", "1")

		c.JSON(http.StatusOK, gin.H{
			"name":name,
			"age":age,
		})
	})
}

//query + post
func paramQueryPost(router *gin.Engine)  {
	router.POST("query_post", func(c *gin.Context) {
		home := c.Query("home")
		city := c.DefaultQuery("city", "太原")

		name := c.PostForm("name")
		age := c.DefaultPostForm("age", "26")

		c.JSON(http.StatusOK, gin.H{
			"home":home,
			"city":city,
			"name":name,
			"age": age,
		})
	})
}

// Map作为Query
// query_map?ids[a]=1&ids[b]=2 使用的是 map
// query_map?ids=1&ids=2 使用的是 array
func paramQueryMap(router *gin.Engine)  {
	router.GET("query_map", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		ids2 := c.QueryArray("ids")


		fmt.Println(ids)

		c.String(http.StatusOK, fmt.Sprintf("map: %v, arr: %v", ids, ids2))
	})
}


// Map作为Query
// query_map?ids[a]=1&ids[b]=2 使用的是 map
// query_map?ids=1&ids=2 使用的是 array
func paramQueryMap2(router *gin.RouterGroup)  {
	router.GET("query_map", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		ids2 := c.QueryArray("ids")

		v := db.Get("name1")

		fmt.Println(ids)

		c.String(http.StatusOK, fmt.Sprintf("map: %v, arr: %v, redis %s", ids, ids2, v))
	})
}