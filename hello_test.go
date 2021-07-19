package hello

import "testing"

func TestHello1(t *testing.T) {
	want := "Hello, world."
	if got := Hello1(); got != want {
		t.Errorf("Hello1() = %q, want %q", got, want)
	}
}

func TestHello2(t *testing.T) {
	want := "Hello, world."
	if got := Hello2(); got != want {
		t.Errorf("Hello2() = %q, want %q", got, want)
	}
}
