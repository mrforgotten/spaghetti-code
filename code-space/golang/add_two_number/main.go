package main

import "fmt"

// [add two number](https://leetcode.com/problems/add-two-numbers)
// question from leet code

type ListNode struct {
	Val  int
	Next *ListNode
}

// modified ListNode is not available in leetcode
func (n ListNode) GetNumber() int {
	ans := 0
	var temp ListNode = n
	mult := 1
	for i := 1; true; i++ {
		if !temp.HasNext() {
			ans += temp.Val * mult
			break
		}
		ans += (temp.Val * mult)
		mult *= 10
		temp = *temp.Next
	}
	return ans
}

// func (n ListNode) GetNumber() int {
// 	if !n.HasNext() {
// 		return n.Val
// 	}
// 	return n.Val + n.Next.GetNumber()*10
// }

func (n ListNode) HasNext() bool {
	return n.Next != nil
}

func (n *ListNode) AddNext(val int) {
	newNode := ListNode{
		val,
		nil,
	}

	if n.HasNext() {
		newNode.Next = n.Next
	}

	n.Next = &newNode
}

func main() {
	var l1 *ListNode = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	var l2 *ListNode = &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	ansNode := addTwoNumbers(l1, l2)

	ansNum := ansNode.GetNumber()
	fmt.Print(ansNum)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num := l1.GetNumber() + l2.GetNumber()

	ans := &ListNode{
		num % 10, nil,
	}

	if num < 10 {
		return ans
	} else {
		num /= 10
	}

	var tempAns *ListNode = ans
	for {
		tempAns.AddNext(num % 10)
		num /= 10
		if num == 0 {
			break
		}
		tempAns = tempAns.Next
	}

	return ans
}

/* old version */

// var exceed = 0
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	l1Temp := l1
// 	l2Temp := l2
// 	var l1Val, l2Val, l3Val int

// 	if l1 != nil {
// 		l1Val = l1Temp.Val
// 	} else {
// 		l1Val = 0
// 	}

// 	if l2 != nil {
// 		l2Val = l2Temp.Val
// 	} else {
// 		l2Val = 0
// 	}

// 	l3Val = l1Val + l2Val + exceed

// 	if l3Val >= 10 {
// 		l3Val = l3Val % 10
// 		exceed = 1
// 	} else {
// 		exceed = 0
// 	}

// 	if l1Temp != nil && l1Temp.Next != nil {
// 		l1Temp = l1Temp.Next
// 	} else {
// 		l1Temp = nil
// 	}
// 	if l2Temp != nil && l2Temp.Next != nil {
// 		l2Temp = l2Temp.Next
// 	} else {
// 		l2Temp = nil
// 	}
// 	if l1Temp != nil || l2Temp != nil || exceed != 0 {
// 		l3 := ListNode{l3Val % 10, addTwoNumbers(l1Temp, l2Temp)}
// 		return &l3
// 	}
// 	l3 := ListNode{l3Val % 10, nil}
// 	return &l3
// }
