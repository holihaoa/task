package main

import (
	"fmt"

	"github.com/holihaoa/task/base1"
	_ "github.com/holihaoa/task/base1"
)

const mainName string = "main"

var mainVar string = getMainVar()

var nums [5]int = [5]int{1, 1, 3, 2, 2}

func init() {
	fmt.Println("main init method invoked")
}

func main() {
	onceNumber := base1.GetOnceNumber(nums)
	fmt.Println("getOnceNumber is ", onceNumber)
}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked!")
	return mainName
}
