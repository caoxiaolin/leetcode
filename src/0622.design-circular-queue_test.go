package leetcode

import (
	"testing"
)

func TestMyCricularQueue(t *testing.T) {
	circularQueue := Constructor622(3)
	res := circularQueue.DeQueue() // 返回false
	if res {
		t.Errorf("\n expect: false\nbut got: %v", res)
	}
	res1 := circularQueue.Rear() // 返回-1
	if res1 != -1 {
		t.Errorf("\n expect: -1\nbut got: %v", res1)
	}
	res1 = circularQueue.Front() // 返回-1
	if res1 != -1 {
		t.Errorf("\n expect: -1\nbut got: %v", res1)
	}
	res = circularQueue.EnQueue(1) //返回true
	if res != true {
		t.Errorf("\n expect: true\nbut got: %v", res)
	}
	res = circularQueue.DeQueue() // 返回true
	if res != true {
		t.Errorf("\n expect: true\nbut got: %v", res)
	}
	res = circularQueue.EnQueue(1) //返回true
	if res != true {
		t.Errorf("\n expect: true\nbut got: %v", res)
	}
	res = circularQueue.EnQueue(2) //返回true
	if res != true {
		t.Errorf("\n expect: true\nbut got: %v", res)
	}
	res = circularQueue.EnQueue(3) //返回true
	if res != true {
		t.Errorf("\n expect: true\nbut got: %v", res)
	}
	res = circularQueue.EnQueue(4) // 返回false,队列已满
	if res {
		t.Errorf("\n expect: false\nbut got: %v", res)
	}
	res1 = circularQueue.Rear() // 返回3
	if res1 != 3 {
		t.Errorf("\n expect: 3\nbut got: %v", res1)
	}
	res = circularQueue.IsFull() // 返回true
	if res != true {
		t.Errorf("\n expect: true\nbut got: %v", res)
	}
	res = circularQueue.DeQueue() // 返回true
	if res != true {
		t.Errorf("\n expect: true\nbut got: %v", res)
	}
	res = circularQueue.EnQueue(4) // 返回true
	if res != true {
		t.Errorf("\n expect: true\nbut got: %v", res)
	}
	res1 = circularQueue.Rear() // 返回4
	if res1 != 4 {
		t.Errorf("\n expect: 4\nbut got: %v", res1)
	}
	res1 = circularQueue.Front() // 返回2
	if res1 != 2 {
		t.Errorf("\n expect: 4\nbut got: %v", res1)
	}
	res = circularQueue.DeQueue() // 返回true
	if res != true {
		t.Errorf("\n expect: true\nbut got: %v", res)
	}
	res = circularQueue.DeQueue() // 返回true
	if res != true {
		t.Errorf("\n expect: true\nbut got: %v", res)
	}
	res = circularQueue.DeQueue() // 返回true
	if res != true {
		t.Errorf("\n expect: true\nbut got: %v", res)
	}
	res = circularQueue.DeQueue() // 返回false
	if res {
		t.Errorf("\n expect: true\nbut got: %v", res)
	}
}
