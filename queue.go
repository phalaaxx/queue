package queue

import (
	"sync"
)

/* Queue represents a strings queue */
type Queue struct {
	Mutex  sync.Mutex
	Queue  []interface{}
	Signal *sync.Cond
}

/* Push element at the end of the queue */
func (q *Queue) Push(x interface{}) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	// make sure x does not exist
	for _, item := range q.Queue {
		if item == x {
			return
		}
	}
	// add new item at end of queue
	q.Queue = append(q.Queue, x)
	q.Signal.Broadcast()
}

/* Pop element from the top of the queue */
func (q *Queue) Pop() interface{} {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	// wait for signal if queue is empty
	if len(q.Queue) == 0 {
		q.Signal.Wait()
	}
	// pop element from beginning of queue
	x := q.Queue[0]
	q.Queue = q.Queue[1:len(q.Queue)]
	return x
}

/* Reset removes all elements currently in queue */
func (q *Queue) Reset() {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	// wait for signal
	if len(q.Queue) == 0 {
		q.Signal.Wait()
	}
	// reset queue and ignore all elements
	q.Queue = []interface{}{}
}

/* NewQueue creates a new message queue object */
func NewQueue() *Queue {
	queue := new(Queue)
	queue.Signal = sync.NewCond(&queue.Mutex)
	return queue
}