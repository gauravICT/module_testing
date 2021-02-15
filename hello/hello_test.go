package hello

import "testing"

//TestHello :
func TestHello(t *testing.T) {
	//want := "Hello, Everyone My Name is Gaurav !"
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

// TestProverb :
func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}
