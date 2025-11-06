package advanced

import (
	"fmt"
	"time"
)

// Stateful goroutines keep track of their previous state (or shared state)
// Generally used for data stream processing and task processing

type StatefulWorker struct {
	count int
	ch chan int
}

func (w *StatefulWorker) Start() {
	go func() {
		for {
			select {
			case value := <- w.ch:
				w.count += value
				fmt.Println("Current count:", w.count)
			}
		}
	}()
}

func (w *StatefulWorker) Send(value int) {
	w.ch <- value
}

func StatefulGoroutinesDemo() {
	worker := &StatefulWorker{
		ch: make(chan int),
	}

	worker.Start()

	for i := range 5 {
		worker.Send(i)
		time.Sleep(500 * time.Millisecond)
	}
}