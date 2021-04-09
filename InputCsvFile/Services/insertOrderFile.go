package Services

import (
	"fmt"
	"time"
)

func insertOrder(ch chan jinchangBox) {

	pool:
	for true {
		select {
		case tBox := <-ch:
			// 插入数据库
			insert_db_box(tBox)

			//fmt.Println(tBox)
			break
		case <-time.After(2 * time.Second): //上面的ch如果一直没数据会阻塞，那么select也会检测其他case条件，检测到后3秒超时
			fmt.Println("超时")
			close(ch)
			break pool
		}
	}

	fmt.Println("exit")
}

