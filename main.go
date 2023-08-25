package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rd := func() interface{} { return rand.Intn(50000000) }

	done := make(chan interface{})
	defer close(done)

	// 4.6.1
	intCh := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intCh, 2), 1), 2)
	for v := range pipeline {
		fmt.Println(v)
	}
	// 4.6.2
	for num := range take(done, repeatFn(done, rd), 5) {
		fmt.Printf("%v\n", num)
	}

	// 4.7
	start := time.Now()
	randIntCh := toInt(done, repeatFn(done, rd))
	fmt.Println("Primes:")
	for prime := range take(done, primeFinder(done, randIntCh), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v\n", time.Since(start))

}
