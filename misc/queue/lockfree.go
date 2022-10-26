package queue

type Node struct {
	data interface{}
	next *Node
}

// LockfreeQueue represents a FIFO structure with operations to enqueue
// and dequeue tasks represented as Request
type LockFreeQueue struct {
	head *Node // where things come out (FIFO)
	tail *Node // where things go in (FIFO)
}

// NewQueue creates and initializes a LockFreeQueue
func NewLockFreeQueue() *LockFreeQueue {
	return &LockFreeQueue{nil, nil}
}

// Enqueue adds a series of Request to the queue
func (queue *LockFreeQueue) Enqueue(data interface{}) {
	// node := NewNode(data)
	// tail := queue.tail
	// unsafe.Pointer()

	// nodePtr := unsafe.Pointer(node)
	// tailPtr := unsafe.Pointer(tail)

	// for !atomic.CompareAndSwapPointer(, nil, node) {

	// }
}

// Dequeue removes a Request from the queue
func (queue *LockFreeQueue) Dequeue() interface{} {
	return nil
}

// NewNode creates a Node struct, capable of carrying any data
func NewNode(data interface{}) *Node {
	return &Node{data, nil}
}
