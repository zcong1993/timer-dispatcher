package tasks

import (
	"container/heap"
)

type TimerTask struct {
	index     int
	timestamp int64
	value     string
}

var sequence = 0

type TimerTasksHeap []*TimerTask

func (td TimerTasksHeap) Len() int {
	return len(td)
}

func (td TimerTasksHeap) Less(i, j int) bool {
	return td[i].timestamp > td[j].timestamp
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
	item.timestamp = timestamp
	item.value = value
	heap.Fix(td, item.index)
}

func NewTask(value string, timestamp int64) *TimerTask {
	return &TimerTask{
		value:     value,
		timestamp: timestamp,
	}
}
