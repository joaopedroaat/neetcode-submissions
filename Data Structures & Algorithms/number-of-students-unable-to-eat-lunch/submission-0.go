type Node struct {
	val  int
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func NewQueue(initialValues ...int) *Queue {
	var head, tail, curr *Node

	if len(initialValues) > 0 {
		for _, value := range initialValues {
			if head == nil {
				head = &Node{
					val: value,
				}
				tail = head
				curr = head
				continue
			}

			curr.next = &Node{
				val: value,
			}
			curr = curr.next
			tail = curr
		}
	}

	return &Queue{
		head,
		tail,
	}
}

func (q *Queue) Enqueue(value int) {
	node := &Node{
		val: value,
	}

	q.tail.next = node
	q.tail = q.tail.next
}

func (q *Queue) Dequeue() *Node {
	dequeued := q.head
	q.head = q.head.next
	return dequeued
}

func countStudents(students []int, sandwiches []int) int {
	queue := NewQueue(students...)

	rejected := 0
	for queue.head != nil && rejected != len(sandwiches) {
		if queue.head.val == sandwiches[0] {
			sandwiches = sandwiches[1:]
			queue.Dequeue()
			rejected = 0
		} else {
			dequeued := queue.Dequeue()
			queue.Enqueue(dequeued.val)
			rejected++
		}
	}

	return len(sandwiches)
}
