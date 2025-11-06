package advanced

import (
	"fmt"
	"sync"
	"time"
)

// A leaky bucket rate limiting algorithm is generally used for quality of service in computer networks to prevent congestion
// It accepts variable load from the requests and ensures that a smooth rate of the incoming requests are processed
// Additional requests are held in the bucket and wait before they are processed
// A steady flow of requests are processed (simulated by a bucket leaking a small portion of water)

type LeakyBucket struct {
	capacity int
	leakRate time.Duration
	tokens int
	lastLeak time.Time
	mu sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens: capacity,
		lastLeak: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	now := time.Now()
	elapsedTime := now.Sub(lb.lastLeak)
	tokensToAdd := int(elapsedTime / lb.leakRate)
	lb.tokens += tokensToAdd

	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}

	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)

	if lb.tokens > 0 {
		lb.tokens--
		return true
	}

	return false
}

func LeakyBucketDemo() {
	leakyBucketInstance := NewLeakyBucket(5, 500 * time.Millisecond)

	for range 10 {
		if leakyBucketInstance.Allow() {
			fmt.Println("Request accepted!")
		} else {
			fmt.Println("Request rejected!")
		}
		time.Sleep(200 * time.Millisecond)
	}
}