package main

var multiply = func(
	done <-chan interface{},
	intCh <-chan int,
	multiplier int,
) <-chan int {
	multipliedCh := make(chan int)
	go func() {
		defer close(multipliedCh)
		for i := range intCh {
			select {
			case <-done:
				return

			case multipliedCh <- i * multiplier:
			}
		}
	}()
	return multipliedCh
}

var add = func(
	done <-chan interface{},
	intCh <-chan int,
	additive int,
) <-chan int {
	addedCh := make(chan int)
	go func() {
		defer close(addedCh)
		for i := range intCh {
			select {
			case <-done:
				return
			case addedCh <- i + additive:
			}
		}
	}()
	return addedCh
}

var isPrimeInefficient = func(done <-chan interface{}, n int) <-chan bool {
	resultCh := make(chan bool)
	go func() {
		defer close(resultCh)
		if n <= 1 {
			resultCh <- false
			return
		}
		for i := 2; i < n; i++ {
			select {
			case <-done:
				return
			default:
				if n%i == 0 {
					resultCh <- false
					return
				}
			}
		}
		resultCh <- true
	}()
	return resultCh
}

var primeFinder = func(done <-chan interface{}, intCh <-chan int) <-chan interface{} {
	primeCh := make(chan interface{})
	go func() {
		defer close(primeCh)
		for i := range intCh {
			ok := make(chan bool, 1)
			select {
			case <-done:
				return
			case ok <- <-isPrimeInefficient(done, i):
				if <-ok {
					primeCh <- i
				}
			}
			close(ok)
		}
	}()

	return primeCh
}
