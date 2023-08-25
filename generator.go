package main

var generator = func(done <-chan interface{}, integers ...int) <-chan int {
	intCh := make(chan int, len(integers))
	go func() {
		defer close(intCh)
		for _, i := range integers {
			select {
			case <-done:
				return

			case intCh <- i:
			}
		}
	}()
	return intCh
}

var take = func(
	done <-chan interface{},
	valueCh <-chan interface{},
	num int,
) <-chan interface{} {
	takeCh := make(chan interface{})
	go func() {
		defer close(takeCh)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeCh <- <-valueCh:
			}
		}
	}()
	return takeCh
}

var repeatFn = func(
	done <-chan interface{},
	fn func() interface{},
) <-chan interface{} {
	valueCh := make(chan interface{})
	go func() {
		defer close(valueCh)
		for {
			select {
			case <-done:
				return
			case valueCh <- fn():
			}
		}
	}()
	return valueCh
}

var toInt = func(done <-chan interface{}, valueCh <-chan interface{}) <-chan int {
	intCh := make(chan int)
	go func() {
		defer close(intCh)
		for i := range valueCh {
			select {
			case <-done:
				return
			case intCh <- i.(int):
			}
		}
	}()
	return intCh
}
