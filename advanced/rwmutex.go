package advanced

import (
	"fmt"
	"sync"
)

// RWMutex stands for read-write mutex
// It is generally used in the case of read heavy and write light operations


var (
	rwmu sync.RWMutex
	counter int
)

func readCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	rwmu.RLock()
	fmt.Println("Read counter:", counter)
	rwmu.RUnlock()
}

func writeCounter(wg *sync.WaitGroup, value int) {
	defer wg.Done()
	rwmu.Lock()
	counter = value
	fmt.Printf("Wrote value of %d into the counter\n", value)
	rwmu.Unlock()
}

func RWMutexDemo() {
	var wg sync.WaitGroup

	for range 5 {
		wg.Add(1)
		go readCounter(&wg)
	}

	wg.Add(1)
	go writeCounter(&wg, 18)
	
	wg.Wait()

}