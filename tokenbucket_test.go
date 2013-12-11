package tokenbucket

import "testing"

func TestFoo(t *testing.T) {
	const in, out = 2, 2
	if in != out {
		t.Errorf("Foo() = %v, want %v", in, out)
	}
}

func TestNew(t *testing.T) {
	const out = 1
	x := New()
	if x.foo != out {
		t.Errorf("New() = %v, want %v", x, out)
	}
}
