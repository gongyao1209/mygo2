package main

import (
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

	paramInPath(r)

	r.Run(":12345")
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