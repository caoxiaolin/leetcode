package leetcode

import (
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	list1 := []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
	listnode1 := InitListNode(list1)

	list2 := []int{13, 20, 45, 47, 52, 88}
	listnode2 := InitListNode(list2)

	res := getIntersectionNode(listnode1, listnode2)
	if res != nil {
		t.Errorf("\n expect: %v\nbut got: %v", nil, res)
	}

	listnode2.Next.Next = listnode1.Next.Next.Next.Next.Next

	res = getIntersectionNode(listnode1, listnode2)
	if res != listnode2.Next.Next {
		t.Errorf("\n expect: %v\nbut got: %v", listnode2.Next.Next, res)
	}

	list1 = []int{13, 20, 45, 47, 52, 88}
	listnode1 = InitListNode(list1)

	list2 = []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
	listnode2 = InitListNode(list2)

	res = getIntersectionNode(listnode1, listnode2)
	if res != nil {
		t.Errorf("\n expect: %v\nbut got: %v", nil, res)
	}
}
