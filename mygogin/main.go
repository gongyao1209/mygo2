package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mygo2/db"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	db.GetData()
}

func main1() {

	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":9091", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	// 写日志
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")

	os.Open()
	//gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.New()

	r.Use(gin.Logger())

	testPost(r)

	r.GET("/ping", func(c *gin.Context) {
		db.GetData()
		//da, _ := json.Marshal(db.GetData())

		c.JSON(200, gin.H{
			"message": "pong1",
			//da
		})
	})

	r.Run(":9091")
}

// =====================================POST传递参数=========================================

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"-"`
}

func testPost(r *gin.Engine) {
	r.POST("/loginJson", testPost1)

	r.POST("/loginPost", func(ctx *gin.Context) {
		var form Login
		if err := ctx.ShouldBind(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
}

func testPost1(ctx *gin.Context) {
	var json Login
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.User != "manu" || json.Password != "123" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
