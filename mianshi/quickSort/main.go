package main

import "fmt"

func main()  {

	var a []int
	a = []int{5, 2, 8, 4, 3, 9, 5, 232, 2,3, 21,3 ,4, 3,32, 43, 5,45,34}

	QuickSort(a)

	fmt.Println(a)
}

func quickSort1(arr []int)  {

	i := 0
	j := len(arr) - 1

	if i == j {
		return
	}

	tmp := arr[i]

	for i < j {
		// 从后往前找 小于 tmp的
		for arr[j] >= tmp && i < j {
			j--
		}
		arr[i] = arr[j]

		for arr[i] < tmp && i < j {
			i++
		}
		arr[j] = arr[i]
	}

	// 一遍完了
	arr[i] = tmp


	fmt.Println(arr)
	return

	//quickSort(arr[0:i])
	//quickSort(arr[i+1:])
}

// 第一种写法
func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for i <= p && values[i] <= temp {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func QuickSort(values []int) {
	if len(values) <= 1 {
		return
	}
	quickSort(values, 0, len(values)-1)
}