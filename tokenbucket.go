package tokenbucket

import (
	"time"
)

type TokenBucket struct {
	rate      int
	size      int
	timestamp int64
}

func New(rate int, size int) *TokenBucket {
	tb := &TokenBucket{
		rate:      rate,
		size:      size,
		timestamp: time.Now().Unix(),
	}
	return tb
}

func (tb *TokenBucket) GetTokens() int {
	now := time.Now().Unix()
	seconds := int(0)
	if now > tb.timestamp {
		seconds = int(now - tb.timestamp)
	}
	return seconds * tb.rate
}
