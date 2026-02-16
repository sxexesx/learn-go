package main

import (
	"fmt"
	"runtime"
	"time"
)

type data struct {
	a []byte
}

func main() {
	t1 := time.NewTicker(100 * time.Microsecond)

	go func() {

		for range t1.C {
			getData()
		}
	}()

	t := time.NewTicker(1 * time.Second)

	var m runtime.MemStats
	now := time.Now()

	for curr := range t.C {
		runtime.ReadMemStats(&m)

		fmt.Printf("GC_enabled %v    GC_runs %v    live_now %v   pause_total_ms %.2f    time %5.0f sec\n",
			m.EnableGC, m.NumGC, m.Mallocs-m.Frees, float64(m.PauseTotalNs)/1000/1000, curr.Sub(now).Seconds())
	}
}

//go:noinline
func getData() *data {
	arr := [300]byte{}

	return &data{
		a: arr[:10],
	}
}
