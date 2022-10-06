package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		name  string
		input int
		want  string
	}{
		{name: "input 1 should be 1", input: 1, want: "1"},
		{name: "input 2 should be 2", input: 2, want: "2"},
		{name: "input 3 should be Fizz", input: 3, want: "Fizz"},
		{name: "input 4 should be 4", input: 4, want: "4"},
		{name: "input 5 should be Buzz", input: 5, want: "Buzz"},
		{name: "input 6 should be Fizz", input: 6, want: "Fizz"},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			got := FizzBuzz(tt.input)

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
