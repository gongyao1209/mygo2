package main

import (
	"mygo2/db"
	"time"

	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main()  {
	// 测试数据库
	db.GetData()
	return

	cur := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(cur)

	cur2 := time.Now()
	ti2, _ := time.Parse("2006-01-02 15:04:05", "2019-01-26 10:06:24")
	a := cur2.Before(ti2)
	fmt.Println(a)
	return

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

		// 开辟go程完成后续工作，这个是没有问题的。工作肯定能做完
		go func() {
			for i := 0; i < 10; i++ {
				time.Sleep(5 * time.Second)
				fmt.Println(time.Now().String(), " check goroutine is worked, i := ", i)
			}
		}()

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
//package main
//
//import (
//	"context"
//	"fmt"
//	"log"
//	"mygo2/httptest"
//	"net/http"
//	"runtime"
//	"sync"
//	"time"
//)
//
//func main()  {
//	http.HandleFunc("/", httptest.SayHello)
//	http.HandleFunc("/login", httptest.Login)
//	http.HandleFunc("/test", httptest.Test)
//
//	err := http.ListenAndServe(":1234", nil)
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
//	//freeWheel.TestReflect()
//}
//
//func worker(ctx context.Context, num int, wg *sync.WaitGroup)  {
//	defer wg.Done()
//
//	for {
//		select {
//		case <-ctx.Done():
//			fmt.Println("quit goroutine", num)
//			return
//		default:
//			fmt.Println("goroutine", num)
//		}
//	}
//}
//
//func QuitGoroutine()  {
//	fmt.Println("goroutine num1  :", runtime.NumGoroutine())
//	ctx, cancel := context.WithCancel(context.Background())
//	var wg sync.WaitGroup
//
//	wg.Add(1)
//	go worker(ctx, 1, &wg)
//
//	wg.Add(1)
//	go worker(ctx, 2, &wg)
//
//	fmt.Println("goroutine num2  :", runtime.NumGoroutine())
//
//	time.Sleep(1 * time.Second)
//	cancel()
//	wg.Wait()
//
//	fmt.Println("goroutine num3  :", runtime.NumGoroutine())
//}
//
//func prin(i, a, b int) int {
//	sum := a + b
//	fmt.Printf("%d) a = %d, b = %d, a + b = %d\n", i, a, b, sum)
//	return sum
//}
//
//func main1()  {
//	a := 1
//	b := 2
//	defer func() {
//		prin(1, a, b)
//	}()
//
//	a = 3
//	b = 4
//	defer func() {
//		prin(2, a, b)
//	}()
//	return
//}
//
//
//func twoSum(nums []int, target int) []int {
//	temp := make(map[int]int)
//	res := make([]int, 2)
//
//	for k , v := range nums {
//		if _, ok := temp[v]; !ok {
//			temp[v] = k
//		}
//	}
//
//	for k , v := range nums {
//		t := target - v
//
//		if a, ok := temp[t]; ok {
//			if k == a {
//				continue
//			}
//
//			res[0] = k
//			res[1] = a
//
//			return res
//		}
//	}
//
//	return res
//}
//
