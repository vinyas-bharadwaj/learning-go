package advanced

import (
	"fmt"
	"time"
)

// Token bucket is a rate limiting algorithm which uses tokens to allow or deny incoming requests
// We define a fixed amount of tokens (allowed requests) which may refill after a set time interval (refillTime)

type RateLimiter struct {
	// We use empty structs since it occupies zero space in memory
	tokens chan struct{}
	refillTime time.Duration
}

func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	r1 := &RateLimiter{
		tokens: make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}
	for range rateLimit {
		r1.tokens <- struct{}{}
	}
	// We will keep refilling the tokens by creating a goroutine that runs in the background
	go r1.startRefill()
	return r1
}

func (rl *RateLimiter) startRefill() {
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()

	for range ticker.C {
		select {
		// We push an empty struct into the channel
		// Since the channel is buffered, this will block if it is full
		case rl.tokens <- struct{}{}:
		default:  // We do not perform any operation in default case
		}
	}
}

func (rl *RateLimiter) allow() bool {
	select {
	// We return true if we receive any values from the tokens channel
	case <- rl.tokens:
		return true
	default:
		return false
	}
}

func TokenBucketDemo() {
	// This rate limiter allows 5 requests
	// The tokens are refreshed every 1 second
	RateLimiter := NewRateLimiter(5, time.Second)

	// We try to simulate 20 requests (Rate limiter allows only 5)
	for range 20 {
		if RateLimiter.allow() {
			fmt.Println("Request allowed!")
		} else {
			fmt.Println("Request denied!")
		}
		time.Sleep(300 * time.Millisecond)
	}
}