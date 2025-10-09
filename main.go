package main

import (
	"fmt"
	"time"

	"github.com/holihaoa/task/base1"
	"github.com/holihaoa/task/base2"
	_ "github.com/holihaoa/task/base2"
)

const mainName string = "main"

var nums [5]int = [5]int{1, 1, 3, 2, 2}

func init() {
}

func main() {
	// 任务一
	//task1()

	// 任务二
	intParam := 23
	intResult := base2.SetIntValue(&intParam) // 指针一任务,通过指针修改int值
	fmt.Println("Pointer change value result is", intResult)

	slideValue := []int{1, 2, 3}
	slideResult := base2.SetIntSlideValue(&slideValue) // 指针二任务,通过指针修改切片值
	fmt.Println("Pointer change slide value result is", slideResult)

	base2.GoroutineTask1() // 协程一任务,分为两个协程分别打印奇偶数
	time.Sleep(time.Second)

	// 协程任务二,任务调度器
	GoroutineT2()

	base2.OopTask1() // 面向对象任务一

	base2.OopTask2() // 面向对象任务二

	base2.ChannelTask1() // 通道任务一

	base2.ChannelTask2() // 通道任务二

	base2.LockTask1() //锁任务一

	base2.LockTask2() // 锁任务二
}

func GoroutineT2() {
	taskQ := base2.TaskQueue{Tasks: make([]base2.Task, 0)}
	taskQ.Tasks = append(taskQ.Tasks, base2.Task{Name: "task1", Work: func() {
		println("work1 is working")
		time.Sleep(100 * time.Millisecond)
	}})
	taskQ.Tasks = append(taskQ.Tasks, base2.Task{Name: "task2", Work: func() {
		println("work2 is working")
		time.Sleep(1 * time.Second)
	}})
	taskQ.Tasks = append(taskQ.Tasks, base2.Task{Name: "task3", Work: func() {
		println("work3 is working")
		time.Sleep(500 * time.Millisecond)
	}})
	timeCostResult := base2.GoroutineTask2(taskQ)
	for _, value := range timeCostResult {
		fmt.Printf("task %s cost result is %s \n", value.TaskName, value.Duration)
	}
}

func task1() {
	// 任务一
	onceNumber := base1.GetOnceNumber(nums) // 第一题 只出现一次的数字
	fmt.Println("getOnceNumber is ", onceNumber)
	ispalindrome := base1.IsPalindrome(0) // 第二题 回文数
	fmt.Println("IsPalindrome is ", ispalindrome)
	isvalide := base1.IsValid("(]") // 第三题 有效括号
	fmt.Println("IsValid is ", isvalide)
	longestCommon := base1.LongestCommonPrefix([]string{"aaa", "aa", "aaa"}) // 第四题 公共前缀
	fmt.Println("LongestCommonPrefix is ", longestCommon)
	digits := base1.PlusOne([]int{0}) // 第五题 加一
	fmt.Println("PlusOne result is ", digits)
	num := base1.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}) // 第六题 去重
	fmt.Println("RemoveDuplicates num is ", num)
	mergeArr := base1.Merge([][]int{{4, 5}, {1, 4}, {0, 1}}) // 第七题 合并区间
	fmt.Println("mergeArr is ", mergeArr)
	twosum := base1.TwoSum([]int{1, 6142, 8192, 10239}, 18431) // 第八题 两数之和
	fmt.Println("twosum is ", twosum)
}
