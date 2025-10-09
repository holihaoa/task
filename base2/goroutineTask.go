package base2

import "fmt"

/*题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。*/

func GoroutineTask1() {
	go func() {
		for i := 1; i < 11; i += 2 {
			fmt.Println("goroutine1 print odd number :", i)
		}
	}()

	go func() {
		for i := 2; i < 11; i += 2 {
			fmt.Println("goroutine2 print even number:", i)
		}
	}()
}
