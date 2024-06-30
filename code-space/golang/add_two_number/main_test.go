package main

import "testing"

func isEqual(args ...any) bool {
	argsLen := len(args)
	for i, v := range args {
		for j := i + 1; j < argsLen-1; j++ {
			if v != args[j] {
				return false
			}
		}
	}

	return true
}
func TestListNode(t *testing.T) {
	var ln *ListNode = &ListNode{9, &ListNode{5, &ListNode{7, nil}}}
	t.Logf("Test ListNode")
	{
		t.Log("ListNode: func GetNumber")
		expected := 759
		actual := ln.GetNumber()
		if !isEqual(expected, actual) {
			t.Errorf("Test GetNumber expected %v actual %v", expected, actual)
		}
	}

	{
		t.Log("ListNode: func HasNext")
		expected := true
		actual := ln.HasNext()
		if !isEqual(expected, actual) {
			t.Errorf("Test HasNext expected %v actual %v", expected, actual)
		}
	}

	{
		t.Log("ListNode: func AddNext")

		addNextVal := 5
		ln.AddNext(addNextVal)

		expected := 7559
		actual := ln.GetNumber()
		if !isEqual(expected, actual) {
			t.Errorf("Test HasNext expected from GetNumber %v actual %v", expected, actual)
		}
	}
}

func TestAddTwoNumbers(t *testing.T) {
	var l1 *ListNode = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	var l2 *ListNode = &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	ansNode := addTwoNumbers(l1, l2)

	actual := ansNode.GetNumber()
	expected := 807

	t.Log("Test AddTwoNumbers")
	if !isEqual(expected, actual) {
		t.Errorf("Test AddTwoNumbers expected from GetNumber %v actual %v", expected, actual)
	}
}
