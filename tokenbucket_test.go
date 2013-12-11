package tokenbucket

import (
	"fmt"
	"testing"
	"time"
)

func TestFoo(t *testing.T) {
	const in, out = 2, 2
	if in != out {
		t.Errorf("Foo() = %v, want %v", in, out)
	}
}

func TestNew(t *testing.T) {
	const rate, size = 1, 2
	tb := New(1, 2)
	if tb.rate != rate {
		t.Errorf("New() = %v, want %v", tb, rate)
	}
	if tb.size != size {
		t.Errorf("New() = %v, want %v", tb, size)
	}
	fmt.Println("time: ", tb.timestamp)
	if tb.timestamp <= 0 {
		t.Errorf("Invalid timestamp: %v", tb.timestamp)
	}
}

func TestGetTokens(t *testing.T) {
	tb := New(1, 1)
	var out = 0
	if x := tb.GetTokens(); x != out {
		t.Errorf("GetTokens() = %v, want %v", x, out)
	}
	time.Sleep(time.Second)
	out = 1
	if x := tb.GetTokens(); x != out {
		t.Errorf("GetTokens() = %v, want %v", x, out)
	}
	time.Sleep(5 * time.Second)
	out = 6
	if x := tb.GetTokens(); x != out {
		t.Errorf("GetTokens() = %v, want %v", x, out)
	}
}
