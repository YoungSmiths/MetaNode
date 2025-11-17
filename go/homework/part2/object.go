package main

import (
	"fmt"
	"math"
)

func main() {
	rectangle := Rectangle{1, 2}

	fmt.Println("Rectangle area is ", rectangle.Area())
	fmt.Println("Rectangle perimeter is ", rectangle.Perimeter())

	circle := Circle{10}
	fmt.Println("Circle area is ", circle.Area())
	fmt.Println("Circle perimeter is ", circle.Perimeter())

	employee := Employee{Person{"Ryan", 18}, 1}
	employee.PrintInfo()
}

/*
*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) Area() float64 {
	fmt.Println("Area of Rectangle")
	fmt.Println("height is", r.height, ", width is", r.width)
	return r.height * r.width
}

func (r Rectangle) Perimeter() float64 {
	fmt.Println("Perimeter of Rectangle")
	fmt.Println("height is", r.height, ", width is", r.width)
	return (r.height + r.width) * 2
}

type Circle struct {
	r float64
}

func (c Circle) Area() float64 {
	fmt.Println("Area of Circle")
	fmt.Println("r is ", c.r)
	return math.Pi * c.r * 2
}

func (c Circle) Perimeter() float64 {
	fmt.Println("Perimeter of Circle")
	fmt.Println("r is ", c.r)
	return math.Pi * c.r * c.r
}

/**
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Println("ID:", e.EmployeeID)
	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)
}
