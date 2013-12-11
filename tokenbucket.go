package tokenbucket

import (
	"time"
)

type TokenBucket struct {
	count     int
	rate      int
	size      int
	timestamp int64
}

// Creates a new TokenBucket.
func New(rate int, size int) *TokenBucket {
	tb := &TokenBucket{
		count:     0,
		rate:      rate,
		size:      size,
		timestamp: time.Now().Unix(),
	}
	return tb
}

// Returns the current count of available tokens.
func (tb *TokenBucket) Tokens() int {
	now := time.Now().Unix()
	seconds := int(0)
	if now > tb.timestamp {
		seconds = int(now - tb.timestamp)
	}
    tb.timestamp = time.Now().Unix()
    tokens := seconds * tb.rate
    if tokens > tb.size {
        return tb.size
    }
	return tokens
}

// Consumes tokens. Returns true if there were sufficient tokens available.
func (tb *TokenBucket) Consume(count int) bool {
    available := tb.Tokens() + tb.count
    if count <= available {
        tb.count += available - count
        return true
    } else {
        return false
    }
}
