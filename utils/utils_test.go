package utils

import (
	"testing"
)

func TestMinInt(t *testing.T) {
	const a, b, out = 1, 2, 1
	if x := MinInt(a, b); x != out {
		t.Errorf("Foo() = %v, want %v", x, out)
	}
}

func TestMinInt2(t *testing.T) {
	const a, b, out = 3000, 100, 100
	if x := MinInt(a, b); x != out {
		t.Errorf("Foo() = %v, want %v", x, out)
	}
}
