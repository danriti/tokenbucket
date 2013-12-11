package tokenbucket

type TokenBucket struct {
	foo int
}

func New() *TokenBucket {
	tb := &TokenBucket{
		foo: 1,
	}
	return tb
}
