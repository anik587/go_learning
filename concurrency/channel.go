package concurrency

import (
	"fmt"
	"sync"
)

func Channel() {
	mychan := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {

		fmt.Println(<-mychan)
		fmt.Println(<-mychan)
		val, isChannelOpen := <-mychan
		fmt.Println(isChannelOpen)
		if isChannelOpen {
			fmt.Println(val)
		}
		wg.Done()
	}(mychan, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		mychan <- 5
		mychan <- 6
		//close(mychan)
		mychan <- 0 //zero is confusing, if channel is closed it will send 0
		wg.Done()
	}(mychan, wg)

	wg.Wait()
}
