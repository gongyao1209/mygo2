package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//计算节点的深度
	temp := l1
	l1_deep := 0
	for temp != nil {
		temp = temp.Next
		l1_deep++
	}

	temp = l2
	l2_deep := 0
	for temp != nil {
		temp = temp.Next
		l2_deep++
	}

	var len_ *ListNode
	var low *ListNode
	if l1_deep >= l2_deep {
		len_ = l1
		low = l2
	} else {
		len_ = l2
		low = l1
	}

	//当前节点
	cur := &ListNode{}
	//下一个节点
	var next *ListNode
	//头节点
	head := cur

	sum := 0 //每一次相加的和

	for len_.Next != nil { //一直到最后一次
		cur.Val = sum

		var temp_sum int
		if low == nil {
			temp_sum = len_.Val
		} else {
			temp_sum = len_.Val + low.Val
		}
		cur.Val += temp_sum

		if cur.Val >= 10 {
			sum = int(cur.Val / 10)
			cur.Val = cur.Val % 10
		} else {
			sum = 0
		}

		len_ = len_.Next
		if low != nil {
			low = low.Next
		}

		next = &ListNode{} //新建一个
		cur.Next = next
		cur = cur.Next
	}

	//temp_sum2 := 0
	if low == nil {
		cur.Val = len_.Val + sum
	} else {
		cur.Val = len_.Val + low.Val + sum
	}

	if cur.Val >= 10 {
		next = &ListNode{}
		next.Val = int(cur.Val / 10)
		cur.Next = next

		cur.Val = cur.Val % 10
	}

	return head
}

/**
 * 给定一个字符串，找出不含有重复字符的最长子串的长度。
 *
 */
func LengthOfLongestSubstring(s string) int {
	len_s := len(s)

	res := make(map[int]([]string))

	for i := 0; i < len_s; i++ {
		for j := i + 1; j <= len_s; j++ {
			res[j-i] = append(res[j-i], s[i:j])
		}
	}

	fmt.Println(res)

	max := len(res)
	for true {
		strs, ok := res[max]
		if !ok {
			max--

			continue
		}
		for _, str := range strs {
			if isRepeat(str) {
				return max
			}
		}

		max--
	}

	return 0
	//return max + 1
}

//判断满足不满足
func isRepeat(s string) bool {
	t_m := make(map[string]int)
	len_s := len(s)

	for i := 0; i < len_s; i++ {
		t_m[string(s[i])] = i
	}

	return len(t_m) == len_s
}
