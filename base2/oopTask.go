package base2

import (
	"fmt"
	"math"
)

/*题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。*/

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func OopTask1() {
	rectangle := Rectangle{Width: 10, Height: 5}
	reectangleArea := rectangle.Area()
	rectanglePerimeter := rectangle.Perimeter()
	fmt.Printf("长方形面积为%s,周长为%s \n", reectangleArea, rectanglePerimeter)
	circle := Circle{Radius: 5}
	circleArea := circle.Area()
	circlePerimeter := circle.Perimeter()
	fmt.Printf("圆形面积为%s,周长为%s \n", circleArea, circlePerimeter)
}

/*
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeId int
}

func (e Employee) PrintInfo() {
	fmt.Printf("当前人物信息名字: %s, 年龄: %d, 工号: %d\n", e.Name, e.Age, e.EmployeeId)
}

func OopTask2() {
	employee := Employee{Person{"Jack", 20}, 30}
	employee.PrintInfo()
}
