package sync

import (
	"sync"
	"testing"
)

type Counter struct {
	val int
	mu  sync.Mutex
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func (c *Counter) Value() int {
	return c.val
}
func NewCounter() *Counter {
	return &Counter{}
}

func TestCounter(t *testing.T) {
	t.Run("increment 3 times", func(t *testing.T) {
		counter := Counter{}

		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Error("value is not 3")
		}

	})

	t.Run("rusn safly concornelty", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		if counter.Value() != wantedCount {
			t.Error("value is not 3")
		}
	})
}
