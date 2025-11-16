package main

import (
	"sync"
	"time"
)

// BaseBenchmark 基准测试基础结构，提供通用的执行逻辑
type BaseBenchmark struct {
	name        string
	description string
	category    string
	testFunc    func(int) // 执行具体测试的函数
	workload    int       // 单个任务的工作量
}

// NewBaseBenchmark 创建基础基准测试
func NewBaseBenchmark(name, description, category string, testFunc func(int), workload int) *BaseBenchmark {
	return &BaseBenchmark{
		name:        name,
		description: description,
		category:    category,
		testFunc:    testFunc,
		workload:    workload,
	}
}

// Name 返回测试名称
func (bb *BaseBenchmark) Name() string {
	return bb.name
}

// Description 返回测试描述
func (bb *BaseBenchmark) Description() string {
	return bb.description
}

// Category 返回测试类别
func (bb *BaseBenchmark) Category() string {
	return bb.category
}

// Run 执行基准测试
func (bb *BaseBenchmark) Run(proc, times int) (res BenchmarkResult) {
	res.Proc = proc
	res.Times = times

	tAll := time.Now()
	defer func() {
		res.Duration = time.Since(tAll)
	}()

	// 顺序执行单核测试，固定5次，剔除最值后求平均
	singleTestCount := 5
	singleTimes := make([]time.Duration, singleTestCount)
	for i := 0; i < singleTestCount; i++ {
		startSingle := time.Now()
		bb.testFunc(bb.workload)
		singleTimes[i] = time.Since(startSingle)
	}
	maxIndex := 0
	minIndex := 0
	for i := 1; i < singleTestCount; i++ {
		if singleTimes[i] > singleTimes[maxIndex] {
			maxIndex = i
		}
		if singleTimes[i] < singleTimes[minIndex] {
			minIndex = i
		}
	}
	// 计算剔除最大值和最小值后的平均值
	var totalSingleTime time.Duration
	averageCount := 0
	for i := 0; i < singleTestCount; i++ {
		if i != maxIndex && i != minIndex {
			totalSingleTime += singleTimes[i]
			averageCount++
		}
	}
	res.SingleDuration = totalSingleTime / time.Duration(averageCount)

	// 多核测试
	p := proc * times
	ch := make(chan time.Duration, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)
	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			bb.testFunc(bb.workload)
			ch <- time.Since(t)
		}()
	}
	wg.Wait()
	res.MultiDuration = time.Since(start) / time.Duration(times)
	close(ch)

	res.Name = bb.Name()
	res.Category = bb.Category()
	res.Ratio = float64(res.MultiDuration / res.SingleDuration)
	res.Score = 0.8*timeToScore(res.SingleDuration) + 0.2*timeToScore(res.MultiDuration)
	return
}
