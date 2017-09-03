package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	// channel을 통해 값을 받아 print
	for i := range ch {
		fmt.Println("Received %d", i)
	}
	wg.Done()  // 작업의 완료 신호를 보냄
}

func main() {
	c := make(chan int)
	go printer(c)
	wg.Add(1)  // 1개의 작업을 WaitGroup에 추가

	// channel을 통해 값 전달
	for i := 1; i <= 10; i++ {
		c <- i
	}

	close(c)  // channel close
	wg.Wait()  // 작업을 기다림
}