package timeTicker

import (
	"testing"
	"time"
)

func TestTimeTicker(t *testing.T) {
	tickNum := 0
	interval := 2
	deadline := 10
	for range TimeTicker(time.Duration(interval)*time.Second, time.Duration(deadline)*time.Second) {
		tickNum++
	}
	if tickNum != deadline/interval {
		t.Errorf("TimeTicker interval %dsecond, deadline %dsencode; want tickNum=%d", interval, deadline, deadline/interval)
	}
}
