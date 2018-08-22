package main

import (
	"fmt"
	"github.com/zcong1993/timer-dispatcher/core/tasks"
	"time"
)

func main() {
	s := tasks.NewScheduler(time.Second*5, 10)

	ch := s.MsgCh()

	println(time.Now().Unix())

	for i := 0; i < 10; i++ {
		s.Push(fmt.Sprintf("haha - %d", i), time.Now().Add(time.Duration(int32(i))*time.Second).Unix())
	}

	println(time.Now().Unix())

	for msg := range ch {
		println(msg)
		println(s.Len())
	}
}
