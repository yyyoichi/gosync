package main

import (
	"math/rand"
	"runtime"
	"testing"
)

var rd = func() interface{} { return rand.Intn(50000000) }

func BenchmarkPipe(b *testing.B) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	randIntCh := toInt(done, repeatFn(done, rd))
	for range take(done, primeFinder(done, randIntCh), 10) {
	}
}

func BenchmarkFunOut(b *testing.B) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	randIntCh := toInt(done, repeatFn(done, rd))

	numFinders := runtime.NumCPU()
	finders := make([]<-chan interface{}, numFinders)
	for i := 0; i < numFinders; i++ {
		// 各 primeFinder ゴルーチンは randIntCh チャネルから受け取る整数を競合なく処理する。
		finders[i] = primeFinder(done, randIntCh)
	}
	for range take(done, funIn(done, finders...), 10) {
	}
}
