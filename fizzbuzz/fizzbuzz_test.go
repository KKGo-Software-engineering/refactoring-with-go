package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	want := "1"

	got := FizzBuzz(1)

	if got != want {
		t.Errorf("FizzBuzz(1) = %q, want %q", got, want)
	}
}
