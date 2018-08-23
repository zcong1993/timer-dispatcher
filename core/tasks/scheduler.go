package tasks

import (
	"container/heap"
	"sync"
	"time"
)

type Scheduler struct {
	interval time.Duration
	ch       chan *TimerTask
	store    *TimerTasksHeap
	stop     chan interface{}
	mu       sync.RWMutex
	busy     bool
}

func NewScheduler(interval time.Duration, chanBufferSize int) *Scheduler {
	hp := &TimerTasksHeap{}
	heap.Init(hp)
	s := &Scheduler{
		interval: interval,
		ch:       make(chan *TimerTask, chanBufferSize),
		store:    hp,
		stop:     make(chan interface{}),
		busy:     false,
	}
	go s.init()
	return s
}

func (s *Scheduler) init() {
	t := time.NewTicker(s.interval)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			if s.busy {
				continue
			}
			s.check()
		case <-s.stop:
			return
		}
	}
}

func (s *Scheduler) check() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.busy = true
	now := time.Now().UnixNano()
	for s.store.Len() > 0 {
		task, ok := heap.Pop(s.store).(*TimerTask)
		if !ok {
			continue
		}
		if task.Timestamp <= now {
			s.ch <- task
		} else {
			// put back task
			s.store.Push(task)
			break
		}
	}
	s.busy = false
}

func (s *Scheduler) Push(value string, timestamp int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t := NewTask(value, timestamp)
	heap.Push(s.store, t)
}

func (s *Scheduler) MsgCh() chan *TimerTask {
	return s.ch
}

func (s *Scheduler) Close() {
	close(s.stop)
}

func (s *Scheduler) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.store.Len()
}
