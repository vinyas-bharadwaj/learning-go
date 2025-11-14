package advanced

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 5

type Buffer struct {
	items []int
	mu sync.Mutex
	cond *sync.Cond
}

func NewBuffer(size int) *Buffer {
	b := &Buffer{
		items: make([]int, 0, size),
	}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *Buffer) Produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// We wait infinitely as long as the size of items is equal to buffer size
	// Wait until buffer is not full
	for len(b.items) == bufferSize {
		b.cond.Wait() // Pauses goroutine until a signal is received
	}

	b.items = append(b.items, item)
	fmt.Printf("Produced: %d; Buffer Size: %d\n", item, len(b.items)) 

	b.cond.Signal() // Signals other wait operations 
}

func (b *Buffer) Consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Wait until buffer has an element in it
	for len(b.items) == 0 {
		b.cond.Wait()
	}

	item := b.items[0]
	b.items = b.items[1:]
	fmt.Printf("Consumed: %d; Buffer Size: %d\n", item, len(b.items))
	b.cond.Signal()
	return item
}

func Producer(b *Buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		b.Produce(i + 100)
		time.Sleep(100 * time.Millisecond)
	}
}

func Consumer(b *Buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 10 {
		b.Consume()
		time.Sleep(200 * time.Millisecond)
	}
}

func SyncNewCondDemo() {
	buffer := NewBuffer(bufferSize)
	var wg sync.WaitGroup

	wg.Add(2)
	go Producer(buffer, &wg)
	go Consumer(buffer, &wg)

	wg.Wait()

}