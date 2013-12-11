package tokenbucket

import (
	"github.com/danriti/tokenbucket/utils"
	"time"
)

type TokenBucket struct {
	available int
	rate      int
	size      int
	timestamp int64
}

// Creates a new TokenBucket.
func New(rate int, size int) *TokenBucket {
	tb := &TokenBucket{
		available: 0,
		rate:      rate,
		size:      size,
		timestamp: time.Now().Unix(),
	}
	return tb
}

// Returns the elapsed time in seconds.
func (tb *TokenBucket) ElapsedTimeSeconds() int {
	now := time.Now().Unix()
	seconds := int(0)

	if now > tb.timestamp {
		seconds = int(now - tb.timestamp)
	}

	return seconds
}

// Returns the current count of available tokens.
func (tb *TokenBucket) Tokens() int {
	tb.UpdateAvailable()
	return tb.available
}

// Update the current count of available tokens.
func (tb *TokenBucket) UpdateAvailable() {
	if tb.available < tb.size {
		seconds := tb.ElapsedTimeSeconds()
		tokens := (tb.rate * seconds) + tb.available
		tb.available = utils.MinInt(tb.size, tokens)
		tb.timestamp = time.Now().Unix()
	}
}

// Consumes tokens. Returns true if there were sufficient tokens available.
func (tb *TokenBucket) Consume(size int) bool {
	if size > tb.Tokens() {
		return false
	}

	tb.available -= size
	return true
}
