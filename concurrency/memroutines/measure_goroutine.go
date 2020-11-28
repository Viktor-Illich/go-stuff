package main

import (
	"fmt"
	"runtime"
	"sync"
)

const numGoroutines = 1e4

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var C <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-C }
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
