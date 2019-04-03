package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mygo2/db"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request)  {
	db.GetData()
}

func main1()  {
	
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":9091", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main()  {
	// 写日志
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)


	r := gin.New()

	r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {
		db.GetData()
		//da, _ := json.Marshal(db.GetData())

		c.JSON(200, gin.H{
			"message":"pong",
			//da
		})
	})

	r.Run(":9091")
}