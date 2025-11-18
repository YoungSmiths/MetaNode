package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var lock sync.Mutex

func main() {
	//add1()
	//add2()
	add3()

}

/*
*
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/
func add1() {
	var index int
	var sw sync.WaitGroup
	sw.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer sw.Done()
			for {
				lock.Lock()
				if index >= 1000 {
					lock.Unlock()
					break
				}
				index++
				fmt.Println(index)
				lock.Unlock()
			}
		}()
	}
	sw.Wait()
	fmt.Println("=== result === ", index)
}

/*
*
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/
func add2() {
	index := int32(0)
	sw := sync.WaitGroup{}
	sw.Add(10)
	for i := 0; i < 10; i++ {
		go func(a *int32) {
			defer sw.Done()
			for k := 0; k < 1000; k++ {
				atomic.AddInt32(a, 1)
			}
		}(&index)
	}
	sw.Wait()
	fmt.Println("=== result === ", index)
}

/*
*
题目 ：启动10个协程，每个协程对计数器进行递增操作，最后输出计数器的值10000。
考察点 ：原子操作、并发数据安全。
*/
func add3() {
	index := int32(0)
	sw := sync.WaitGroup{}
	sw.Add(10)
	for i := 0; i < 10; i++ {
		go func(a *int32) {
			defer sw.Done()
			for {
				currentValue := atomic.LoadInt32(a)
				if currentValue >= 10000 {
					break
				}
				if atomic.CompareAndSwapInt32(a, currentValue, currentValue+1) {
					//if currentValue+1 > 10000 {
					//	break
					//}
				}

			}
		}(&index)
	}
	sw.Wait()
	fmt.Println("=== result === ", index)
}

func add4() {
	index := int32(0)
	sw := sync.WaitGroup{}
	sw.Add(10)
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			defer sw.Done()
			for {
				mu.Lock()
				if index >= 1000 {
					mu.Unlock()
					break
				}
				index++
				mu.Unlock()
			}
		}()
	}

	sw.Wait()
	fmt.Println("=== result === ", index)
}
