package main

import (
	"fmt"
	"github.com/zcong1993/timer-dispatcher/core/tasks"
	"time"
)

func main() {
	s := tasks.NewScheduler(time.Second*5, 10)

	ch := s.MsgCh()

	for i := 0; i < 200; i++ {
		s.Push(fmt.Sprintf("haha - %d", i), time.Now().Add(time.Duration(int32(i))*time.Second).UnixNano())
	}

	for msg := range ch {
		println(msg.Timestamp, msg.Value)
	}
}
