package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("sleep before every print", func(t *testing.T) {
		spyCountdownOperations := &SpyCountdownOperations{}
		Countdown(spyCountdownOperations, spyCountdownOperations)

		want := []string{
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if reflect.DeepEqual(spyCountdownOperations.Calls, want) {
			t.Errorf("got %q, want %q", spyCountdownOperations.Calls, want)
		}
	})

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		ds := &DefaultSleeper{}
		Countdown(buffer, ds)

		got := buffer.String()
		want := `3
2
1
GO!`
		if got != want {
			t.Errorf("got %q, but want %q", got, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
