package utils

import (
	"time"
	"github.com/k0kubun/pp"
)

func Measure(fn func()) {
	start := time.Now()

	fn()

	elapsed := time.Since(start)

	pp.Printf("Execution took %s ms\n", elapsed.Nanoseconds() / int64(time.Millisecond))
}