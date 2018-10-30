package httptest

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func SayHello(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm() //解析参数

	fmt.Println(time.Now().String())
	fmt.Println(r.Form)
	fmt.Println("path: ", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello World") //这个写入到w的是输出到客户端的
}

func Login(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm() //需要解析参数
	fmt.Println("Method: ", r.Method)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Print(t.Execute(w, nil))
	} else {
		// 校验必传
		if len(r.Form["username"]) == 0 || len(r.Form["username"][0]) == 0 {
			fmt.Fprintf(w, "username is need")
			return
		}

		if len(r.Form.Get("age")) == 0 {
			fmt.Fprintf(w, "age is need")
			return
		}
		// 验证是不是数字
		age, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			fmt.Fprintf(w, "age is int")
			return
		}

		if is_han, err := regexp.MatchString("^\\p{Han}+$", r.Form.Get("username")); err == nil {
			if !is_han {
				//fmt.Fprintf(w, "username得是汉字")
				//return
			}
		}

		fmt.Println("username get: ", r.Form.Get("username")) //仅仅获取一个
		fmt.Println("username: ", r.Form["username"]) //和上面的区别就是，下面这个获取slice 是个数组
		fmt.Println("password: ", r.Form["password"])
		fmt.Println("age: ", age)

		template.HTMLEscape(w, []byte(r.Form.Get("username")))
		//fmt.Fprintln(w, r.Form.Get("username"))
	}
}

func Test(w http.ResponseWriter, r *http.Request)  {
	r1 := url.Values{}

	r1.Set("name1", "gongyao1")
	r1.Set("name2", "gongyao2")
	r1.Add("name3", "gongyao3")

	fmt.Fprintf(w, r1.Encode())
}