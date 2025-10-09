package base2

import (
	"fmt"
	"sync"
	"time"
)

/*题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。*/
// 定义任务类型
type Task struct {
	Name string
	Work func()
}

// 定义任务队列
type TaskQueue struct {
	Tasks []Task
	wg    sync.WaitGroup
}

// 定义任务结果
type TaskResult struct {
	TaskName  string
	Duration  time.Duration
	Starttime time.Time
	Endtime   time.Time
}

var resultCh = make(chan TaskResult)

var timeCostResult = make([]TaskResult, 0)

// 任务调度器
func GoroutineTask2(taskQueue TaskQueue) []TaskResult {
	for _, i2 := range taskQueue.Tasks {
		taskResult := TaskResult{}
		taskQueue.wg.Add(1)
		go func(task Task) {
			defer taskQueue.wg.Done()
			taskResult.TaskName = task.Name
			taskResult.Starttime = time.Now()
			task.Work()
			taskResult.Endtime = time.Now()
			taskResult.Duration = taskResult.Endtime.Sub(taskResult.Starttime)
			resultCh <- taskResult
			fmt.Printf("任务 %s 消耗时间 %s \n", task.Name, taskResult.Endtime.Sub(taskResult.Starttime))
		}(i2)
	}

	var collectorWg sync.WaitGroup
	collectorWg.Add(1)

	go func(ch <-chan TaskResult) {
		defer collectorWg.Done()
		for taskResult := range ch {
			timeCostResult = append(timeCostResult, taskResult)
		}
	}(resultCh)
	taskQueue.wg.Wait()
	close(resultCh)
	collectorWg.Wait()
	return timeCostResult
}
