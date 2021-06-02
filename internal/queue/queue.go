package queue

import "github.com/onrcayci/goshell/internal/pcb"

type Queue struct {
	Head *pcb.PCB
	Tail *pcb.PCB
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Push(p *pcb.PCB) {
	if q.Head == nil && q.Tail == nil {
		q.Head = p
		q.Tail = p
	} else {
		oldTail := q.Tail
		oldTail.Next = p
		q.Tail = p
	}
}

func (q *Queue) Pop() *pcb.PCB {
	if q.Tail == nil {
		return nil
	}
	oldHead := q.Head
	q.Head = oldHead.Next
	oldHead.Next = nil
	return oldHead
}
