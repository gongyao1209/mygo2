package panicTest

import "fmt"

func Test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println(err)
			}
		}()
		panic("error, gongyao2")
	}()

	defer func() {
		panic("error, gongyao1")
	}()

	panic("error, gongyao")
}
