package list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	//找到返回的list 以及操作的list
	var r_l, deal_l, cur_node *ListNode
	if l1.Val <= l2.Val {
		r_l = l1
		deal_l = l2
	} else {
		r_l = l2
		deal_l = l1
	}

	cur_node = r_l
	for cur_node.Next != nil {
		if deal_l == nil {
			break
		}

		if cur_node.Val <= deal_l.Val && cur_node.Next.Val > deal_l.Val {
			var temp *ListNode
			deal_l, temp = pop(deal_l)
			temp.Next = cur_node.Next
			cur_node.Next = temp
		}

		cur_node = cur_node.Next
	}

	//如果处理的节点还有的话
	for deal_l != nil {
		var temp *ListNode
		deal_l, temp = pop(deal_l)
		cur_node.Next = temp
		cur_node = cur_node.Next
	}

	return r_l
}

func pop(l2 *ListNode) (*ListNode, *ListNode) {
	var r *ListNode
	if l2 == nil {
		return l2, r
	}

	r = l2
	l2 = r.Next
	r.Next = nil

	return l2, r
}

func getList(n int) *ListNode {
	//temp := listNode
	var l *ListNode
	var pre *ListNode

	for i := 0; i < n; i++ {
		temp := ListNode{Val: i, Next: nil}
		if i == 0 {
			l = &temp
			pre = &temp
		} else {
			pre.Next = &temp
			pre = pre.Next
		}
	}

	return l
}

func TestMergeTwoLists() {
	l1 := getList(4)
	l2 := getList(5)
	print_l(l1)
	fmt.Println("--------")
	print_l(l2)
	fmt.Println("--------")
	//var temp *ListNode
	//l1, temp = pop(l1)
	//fmt.Println(pop(l1))
	l3 := mergeTwoLists(l1, l2)
	print_l(l3)
}

func print_l(l *ListNode) {
	cur := l
	//如果处理的节点还有的话
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
