package leetcode

/**
 * @title 设计循环队列
 *
 * 设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。
 * 循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。
 * 但是使用循环队列，我们能使用这些空间去存储新的值。
 *
 * 你的实现应该支持如下操作：
 * MyCircularQueue(k): 构造器，设置队列长度为 k 。
 * Front: 从队首获取元素。如果队列为空，返回 -1 。
 * Rear: 获取队尾元素。如果队列为空，返回 -1 。
 * enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
 * deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
 * isEmpty(): 检查循环队列是否为空。
 * isFull(): 检查循环队列是否已满。
 *
 * 示例：
 * MyCircularQueue circularQueue = new MycircularQueue(3); // 设置长度为3
 * circularQueue.enQueue(1);  // 返回true
 * circularQueue.enQueue(2);  // 返回true
 * circularQueue.enQueue(3);  // 返回true
 * circularQueue.enQueue(4);  // 返回false,队列已满
 * circularQueue.Rear();  // 返回3
 * circularQueue.isFull();  // 返回true
 * circularQueue.deQueue();  // 返回true
 * circularQueue.enQueue(4);  // 返回true
 * circularQueue.Rear();  // 返回4
 *
 * 提示：
 * 所有的值都在 1 至 1000 的范围内；
 * 操作数将在 1 至 1000 的范围内；
 * 请不要使用内置的队列库。
 *
 * @see https://leetcode-cn.com/problems/design-circular-queue/description/
 * @difficulty Easy
 */

type MyCircularQueue struct {
	Data []int
	Size int
	Head int
	Tail int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor622(k int) MyCircularQueue {
	return MyCircularQueue{make([]int, k), k, -1, -1}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.Head == -1 {
		this.Head = 0
	}
	if this.Tail == this.Size-1 {
		this.Tail = 0
	} else {
		this.Tail++
	}
	this.Data[this.Tail] = value
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.Data[this.Head] = 0
	if this.Head == this.Tail { //只有一个元素，队列置空
		this.Head = -1
		this.Tail = -1
	} else if this.Head == this.Size-1 { //head指向末尾，重头开始
		this.Head = 0
	} else {
		this.Head++
	}
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.Head > -1 {
		return this.Data[this.Head]
	}
	return -1
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.Tail > -1 {
		return this.Data[this.Tail]
	}
	return -1
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if this.Tail == -1 {
		return true
	} else {
		return false
	}
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if this.Tail-this.Head == this.Size-1 || this.Head-this.Tail == 1 {
		return true
	} else {
		return false
	}
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
