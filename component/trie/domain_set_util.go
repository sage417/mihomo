package trie

type RingQueue[T any] struct {
	buf         []T
	head        int
	tail        int
	size        int
	expandTimes int
	capacity    int
}

func NewRingQueue[T any](initCap int) *RingQueue[T] {
	return &RingQueue[T]{
		buf:      make([]T, initCap),
		capacity: initCap,
	}
}

func (q *RingQueue[T]) expand() {
	newCap := q.capacity * 3 / 2
	newBuf := make([]T, newCap)

	if q.head < q.tail {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		copy(newBuf, q.buf[q.head:])
		copy(newBuf[len(q.buf)-q.head:], q.buf[:q.tail])
	}

	q.buf = newBuf
	q.head = 0
	q.tail = q.size
	q.capacity = newCap
	q.expandTimes += 1
}

func (q *RingQueue[T]) Push(v T) bool {
	if q.size >= q.capacity {
		q.expand()
	}
	q.buf[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	q.size++
	return true
}

func (q *RingQueue[T]) Pop() (T, bool) {
	if q.size <= 0 {
		var zero T
		return zero, false
	}
	val := q.buf[q.head]
	var zero T
	q.buf[q.head] = zero

	q.head = (q.head + 1) % q.capacity
	q.size--
	return val, true
}

func (q *RingQueue[T]) Len() int { return q.size }

func (q *RingQueue[T]) Cap() int { return q.capacity }
