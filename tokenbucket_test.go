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

func TestTokens(t *testing.T) {
	tb := New(1, 1)
	var out = 0
	if x := tb.Tokens(); x != out {
		t.Errorf("Tokens() = %v, want %v", x, out)
	}
	time.Sleep(1 * time.Second)
	out = 1
	if x := tb.Tokens(); x != out {
		t.Errorf("Tokens() = %v, want %v", x, out)
	}
	time.Sleep(5 * time.Second)
	out = 1
	if x := tb.Tokens(); x != out {
		t.Errorf("Tokens() = %v, want %v", x, out)
	}
}

func TestTokens2(t *testing.T) {
	tb := New(1, 5)
	var out = 0
	if x := tb.Tokens(); x != out {
		t.Errorf("Tokens() = %v, want %v", x, out)
	}
	time.Sleep(1 * time.Second)
	out = 1
	if x := tb.Tokens(); x != out {
		t.Errorf("Tokens() = %v, want %v", x, out)
	}
	time.Sleep(5 * time.Second)
	out = 5
	if x := tb.Tokens(); x != out {
		t.Errorf("Tokens() = %v, want %v", x, out)
	}
}

func TestConsume(t *testing.T) {
    tb := New(1, 20)
    if tb.Consume(1) != false {
		t.Errorf("Consume() should fail.")
    }
    time.Sleep(1 * time.Second) // 1 token available
    if tb.Consume(1) != true {
		t.Errorf("Consume() should succeed")
    }
    time.Sleep(10 * time.Second) // 10 tokens available
    if tb.Consume(11) != false {
		t.Errorf("Consume() should fail")
    }
}

func TestConsume2(t *testing.T) {
    tb := New(1, 20)
    time.Sleep(5 * time.Second) // 5 tokens available
    if tb.Consume(4) != true {
		t.Errorf("Consume() should succeed.")
    }
    time.Sleep(1 * time.Second) // 2 tokens available
    if tb.Consume(2) != true {
		t.Errorf("Consume() should succeed")
    }
}
