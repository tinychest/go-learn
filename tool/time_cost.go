package tool

import (
	"fmt"
	"time"
)

func TimeCost(f func(), pb ...bool) (d time.Duration) {
	p := append(pb, true)[0]

	s := time.Now()
	f()
	d = time.Now().Sub(s)

	if p {
		fmt.Println(d)
	}
	return
}
