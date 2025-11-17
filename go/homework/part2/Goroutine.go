package main

import (
	"fmt"
	"time"
)

func main() {
	print()
	fmt.Println("start")
	time.Sleep(5 * time.Second)
}

/*
*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/
func print() {
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()
	go print2()

	taskScheduler([]func(){
		func() {
			time.Sleep(1 * time.Second)
			fmt.Println("任务1")
		},
		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("任务2")
		},
		func() {
			fmt.Println("任务3")
		},
	})
}

func print2() {
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}

/*
*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
func taskScheduler(tasks []func()) {
	for _, task := range tasks {
		go func(t func()) {
			start := time.Now()
			t()
			fmt.Printf("任务执行时间：%v\n", time.Since(start))
		}(task)
	}
}
