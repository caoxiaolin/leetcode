package leetcode

import (
    "testing"
)

func TestDetectCycle(t *testing.T) {
	list := []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
    listnode := InitListNode(list)

	res := detectCycle(listnode)

    if res != nil{
        t.Errorf("\n expect: nil\nbut got: %v", res)
    }

    listnode.Next.Next = listnode
    res = detectCycle(listnode)
    if res != listnode.Next.Next{
        t.Errorf("\n expect: %v\nbut got: %v", listnode.Next.Next, res)
    }
}
