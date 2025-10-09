package base2

import (
	"fmt"
	"sync"
)

/*题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。*/

func ChannelTask1() {
	var channel = make(chan int)
	var awg sync.WaitGroup
	var bwg sync.WaitGroup
	awg.Add(1)

	go func() {
		defer awg.Done()
		for num := range channel {
			fmt.Printf("打印数字: %d \n", num)
		}
	}()

	for i := 1; i < 11; i++ {
		bwg.Add(1)
		go func(num int) {
			defer bwg.Done()
			channel <- num
		}(i)
	}
	bwg.Wait()
	close(channel)
	awg.Wait()
}

/*题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。*/

func ChannelTask2() {
	var channel = make(chan int, 100)
	var awg sync.WaitGroup
	var bwg sync.WaitGroup
	awg.Add(1)
	go func() {
		defer awg.Done()
		for num := range channel {
			fmt.Printf("打印数字: %d \n", num)
		}
	}()

	for i := 1; i < 101; i++ {
		bwg.Add(1)
		go func(num int) {
			defer bwg.Done()
			channel <- num
		}(i)
	}
	bwg.Wait()
	close(channel)
	awg.Wait()
}
