package time

import (
	"testing"
	"time"
)

func TestTimeCost(t *testing.T) {
	start := time.Now()

	time.Sleep(2 * time.Second)

	cost := time.Now().Sub(start).String()

	println(cost)
}
