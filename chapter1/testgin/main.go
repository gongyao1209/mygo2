package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()

	g1 := r.Group("/v1")

	g1.GET("name", func(c *gin.Context) {
		c.String(200, "gongyao")
	})

	r.Run(":8001")
}