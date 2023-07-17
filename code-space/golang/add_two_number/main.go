package main

import "fmt"

// [add two number](https://leetcode.com/problems/add-two-numbers)
// question from leet code
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 ListNode = ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	var l2 ListNode = ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	var ans = addTwoNumbers(&l1, &l2)
	var spareAns = *ans

	var ansArr = make([]int, 0)

	for spareAns.Next != nil {
		ansArr = append(ansArr, spareAns.Val)
		spareAns = *spareAns.Next
	}
	ansArr = append(ansArr, spareAns.Val)

	fmt.Println(ansArr)
}

var exceed = 0

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Temp := l1
	l2Temp := l2
	var l1Val, l2Val, l3Val int

	if l1 != nil {
		l1Val = l1Temp.Val
	} else {
		l1Val = 0
	}

	if l2 != nil {
		l2Val = l2Temp.Val
	} else {
		l2Val = 0
	}

	l3Val = l1Val + l2Val + exceed

	if l3Val >= 10 {
		l3Val = l3Val % 10
		exceed = 1
	} else {
		exceed = 0
	}

	if l1Temp != nil && l1Temp.Next != nil {
		l1Temp = l1Temp.Next
	} else {
		l1Temp = nil
	}
	if l2Temp != nil && l2Temp.Next != nil {
		l2Temp = l2Temp.Next
	} else {
		l2Temp = nil
	}
	if l1Temp != nil || l2Temp != nil || exceed != 0 {
		l3 := ListNode{l3Val % 10, addTwoNumbers(l1Temp, l2Temp)}
		return &l3
	}
	l3 := ListNode{l3Val % 10, nil}
	return &l3
}
