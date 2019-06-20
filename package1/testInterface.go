package package1

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type (
	push_driver chan interface{}
	filt        func(v interface{}) bool
)

//
type Pusher struct {
	m       sync.RWMutex
	buffer  int
	timeout time.Duration
	drivers []push_driver
	quit    chan bool
}

func p(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int ", v.(int))
	case string:
		fmt.Println("string ", v.(string))
	case Pusher:
		fmt.Println("Pusher ", v.(Pusher))
	case time.Duration:
		fmt.Println("time ", v)
	default:
		fmt.Println("other ", v)
	}
}

func IsShanxi(filt2 filt, v interface{}) string {
	switch v.(type) {
	case string:
		if filt2 != nil && filt2(v) {
			return "Yes"
		}
		return "OH, No"
	default:
		return "No"
	}
}

func Test101202() {

	//r := rand.New(rand.NewSource(time.Now().Unix()))
	//
	//sl := [10]int{}
	//for i := 0; i < 10; i++ {
	//	sl[i] = r.Intn(1000)
	//	//sl = append(sl, r.Intn(1000))
	//}
	//
	//fmt.Println(len(sl))
	fmt.Println(strings.Replace("gongyao", "yao", "wanghui", 3))

	str := "gongwanghui"

	fmt.Println(strings.Index(str[5:], "yao"))
	os.Exit(1)

	a := IsShanxi(func(v interface{}) bool {
		if v == nil {
			return false
		}

		i := strings.Compare(v.(string), "山西")
		return i == 0
		//return strings.Contains(v.(string), "山西")
	}, "山西")

	fmt.Println(a)
}
