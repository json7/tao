package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/json7/stw"
)

func main() {
	wg := &sync.WaitGroup{}
	wheel := stw.NewTimingWheel(context.TODO())
	timerID := wheel.AddTimer(
		time.Now().Add(2*time.Second),
		1*time.Second,
		stw.NewOnTimeOut(context.TODO(), func(t time.Time, c stw.WriteCloser) { fmt.Printf("TIME OUT AT %s\n", t) }))
	fmt.Printf("Add timer %d\n", timerID)

	wg.Add(1)
	go func() {
		for i := 0; i < 20; i++ {
			select {
			case timeout := <-wheel.TimeOutChannel():
				timeout.Callback(time.Now(), nil)
			}
		}
		wg.Done()
	}()
	wg.Wait()

	wheel.Stop()
}
