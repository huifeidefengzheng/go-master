package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task 任务结构体
type Task struct {
	ID       int
	Name     string
	Function func() error
}

// TaskResult 任务执行结果
type TaskResult struct {
	TaskID   int
	TaskName string
	Duration time.Duration
	Error    error
}

// TaskScheduler 任务调度器
type TaskScheduler struct {
	tasks   []Task
	results chan TaskResult
	wg      sync.WaitGroup
}

// NewTaskScheduler 创建新的任务调度器
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks:   make([]Task, 0),
		results: make(chan TaskResult, 100), // 缓冲通道存储结果
	}
}

// AddTask 添加任务
func (ts *TaskScheduler) AddTask(id int, name string, fn func() error) {
	task := Task{
		ID:       id,
		Name:     name,
		Function: fn,
	}
	ts.tasks = append(ts.tasks, task)
}

// Run 执行所有任务
func (ts *TaskScheduler) Run() []TaskResult {
	// 启动结果收集协程
	ts.wg.Add(1)
	go ts.collectResults()

	// 并发执行所有任务
	for _, task := range ts.tasks {
		ts.wg.Add(1)
		go ts.executeTask(task)
	}

	// 等待所有任务完成
	ts.wg.Wait()

	// 关闭结果通道
	close(ts.results)

	// 收集所有结果
	var results []TaskResult
	for result := range ts.results {
		results = append(results, result)
	}

	return results
}

// executeTask 执行单个任务
func (ts *TaskScheduler) executeTask(task Task) {
	defer ts.wg.Done()

	start := time.Now()

	// 执行任务函数
	err := task.Function()

	duration := time.Since(start)

	// 发送结果到通道
	result := TaskResult{
		TaskID:   task.ID,
		TaskName: task.Name,
		Duration: duration,
		Error:    err,
	}

	ts.results <- result
}

// collectResults 收集结果的协程（示例用途）
func (ts *TaskScheduler) collectResults() {
	defer ts.wg.Done()
	// 这里可以添加结果收集的额外逻辑
}

// 模拟任务函数
func createTaskFunction(taskName string, maxDuration time.Duration) func() error {
	return func() error {
		// 模拟任务执行时间
		duration := time.Duration(rand.Int63n(int64(maxDuration)))
		time.Sleep(duration)

		// 模拟偶尔出现错误
		if rand.Float32() < 0.1 { // 10% 概率出错
			return fmt.Errorf("task %s failed", taskName)
		}

		return nil
	}
}

func main() {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 创建任务调度器
	scheduler := NewTaskScheduler()

	// 添加任务
	tasks := []struct {
		id   int
		name string
		max  time.Duration
	}{
		{1, "数据处理", 500 * time.Millisecond},
		{2, "文件上传", 800 * time.Millisecond},
		{3, "邮件发送", 300 * time.Millisecond},
		{4, "数据库查询", 600 * time.Millisecond},
		{5, "图像处理", 1000 * time.Millisecond},
		{6, "日志分析", 700 * time.Millisecond},
		{7, "缓存清理", 400 * time.Millisecond},
		{8, "报表生成", 900 * time.Millisecond},
	}

	for _, t := range tasks {
		scheduler.AddTask(t.id, t.name, createTaskFunction(t.name, t.max))
	}

	fmt.Println("开始执行任务...")
	start := time.Now()

	// 执行所有任务
	results := scheduler.Run()

	totalDuration := time.Since(start)

	// 打印结果
	fmt.Println("\n任务执行结果:")
	fmt.Println("=====================================")
	for _, result := range results {
		status := "成功"
		if result.Error != nil {
			status = fmt.Sprintf("失败: %v", result.Error)
		}
		fmt.Printf("任务ID: %d | 名称: %-10s | 耗时: %v | 状态: %s\n",
			result.TaskID, result.TaskName, result.Duration, status)
	}

	fmt.Println("=====================================")
	fmt.Printf("总执行时间: %v\n", totalDuration)
	fmt.Printf("完成任务数: %d\n", len(results))
}
