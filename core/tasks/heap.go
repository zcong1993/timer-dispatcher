package tasks

import (
	"container/heap"
)

// TimerTask is struct of deferred task
type TimerTask struct {
	index     int
	Timestamp int64
	Value     string
}

var sequence = 0

// TimerTasksHeap is TimerTasks heap
type TimerTasksHeap []*TimerTask

// Len implement heap Len method
func (td TimerTasksHeap) Len() int {
	return len(td)
}

// Less implement heap Less method
func (td TimerTasksHeap) Less(i, j int) bool {
	return td[i].Timestamp < td[j].Timestamp
}

// Swap implement heap Swap method
func (td TimerTasksHeap) Swap(i, j int) {
	td[i], td[j] = td[j], td[i]
	td[i].index = i
	td[j].index = j
}

// Push implement heap Push method
func (td *TimerTasksHeap) Push(x interface{}) {
	item := x.(*TimerTask)
	item.index = sequence
	sequence++

	*td = append(*td, item)
}

// Pop implement heap Pop method
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

// NewTask return a new timer task
func NewTask(value string, timestamp int64) *TimerTask {
	return &TimerTask{
		Value:     value,
		Timestamp: timestamp,
	}
}
