package tasks

import (
	"container/heap"
)

type TimerTask struct {
	index     int
	Timestamp int64
	Value     string
}

var sequence = 0

type TimerTasksHeap []*TimerTask

func (td TimerTasksHeap) Len() int {
	return len(td)
}

func (td TimerTasksHeap) Less(i, j int) bool {
	return td[i].Timestamp < td[j].Timestamp
}

func (td TimerTasksHeap) Swap(i, j int) {
	td[i], td[j] = td[j], td[i]
	td[i].index = i
	td[j].index = j
}

func (td *TimerTasksHeap) Push(x interface{}) {
	item := x.(*TimerTask)
	item.index = sequence
	sequence++

	*td = append(*td, item)
}

func (td *TimerTasksHeap) Pop() interface{} {
	old := *td
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*td = old[0 : n-1]

	return item
}

func (td *TimerTasksHeap) update(item *TimerTask, value string, timestamp int64) {
	item.Timestamp = timestamp
	item.Value = value
	heap.Fix(td, item.index)
}

func NewTask(value string, timestamp int64) *TimerTask {
	return &TimerTask{
		Value:     value,
		Timestamp: timestamp,
	}
}
