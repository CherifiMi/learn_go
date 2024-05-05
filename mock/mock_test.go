package mock

import (
	"bytes"
	"testing"
)

type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountDown(t *testing.T) {
	assert := func(t testing.TB, got, want any) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, but wanted %q", got, want)
		}
	}

	buffer := &bytes.Buffer{}
	sleeper := SpySleeper{}

	CountDown(buffer, sleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	assert(t, got, want)
	assert(t, sleeper.Calls, 3)
}
