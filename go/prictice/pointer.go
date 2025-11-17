package main

import "fmt"

func main() {
	pointer1()
	pointer2()
}

func pointer1() {
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%s/n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}
}

func pointer2() {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["测试"] = 100
	fmt.Println(b)
}
