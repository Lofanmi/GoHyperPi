package main

import (
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"
)

// SortingBenchmark 排序算法性能测试
type SortingBenchmark struct{}

// NewSortingBenchmark 创建排序测试实例
func NewSortingBenchmark() *SortingBenchmark {
	return &SortingBenchmark{}
}

func (sb *SortingBenchmark) Name() string {
	return "排序算法测试（Sorting Algorithms）"
}

func (sb *SortingBenchmark) Description() string {
	return "测试各种排序算法性能"
}

func (sb *SortingBenchmark) Category() string {
	return "算法性能"
}

func (sb *SortingBenchmark) Run(proc, times int) BenchmarkResult {
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)
	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			sortingTest(100000) // 排序100000个元素
			ch <- time.Since(t).Seconds()
		}()
	}
	wg.Wait()
	duration := time.Since(start)
	close(ch)
	single := 0.0
	for s := range ch {
		single += s
	}
	single = single / float64(p)
	t1 := 10000.0 / single                      // 单核排序速率（elements/s）
	tn := float64(p*10000) / duration.Seconds() // 多核排序速率（elements/s）
	efficiency := tn / t1 / float64(proc)       // 多核效率
	return BenchmarkResult{
		Name:       sb.Name(),
		Category:   sb.Category(),
		Score:      tn / 1000.0, // 归一化得分
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
	}
}

func sortingTest(size int) {
	// 随机数据排序
	data := make([]int, size)
	r := rand.New(rand.NewSource(42))
	for i := 0; i < size; i++ {
		data[i] = r.Intn(size * 10)
	}
	sort.Ints(data)
	// 已排序数据
	sort.Ints(data)
	// 逆序数据
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	sort.Ints(data)
	// 部分排序数据
	for i := size / 2; i < size; i++ {
		data[i] = r.Intn(size * 10)
	}
	sort.Ints(data)
}

// StringBenchmark 字符串处理测试
type StringBenchmark struct{}

func NewStringBenchmark() *StringBenchmark {
	return &StringBenchmark{}
}

func (sb *StringBenchmark) Name() string {
	return "字符串处理（String Processing）"
}

func (sb *StringBenchmark) Description() string {
	return "测试字符串操作性能"
}

func (sb *StringBenchmark) Category() string {
	return "算法性能"
}

func (sb *StringBenchmark) Run(proc, times int) BenchmarkResult {
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)
	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			stringTest(1000)
			ch <- time.Since(t).Seconds()
		}()
	}
	wg.Wait()
	duration := time.Since(start)
	close(ch)
	single := 0.0
	for s := range ch {
		single += s
	}
	single = single / float64(p)
	t1 := 1000.0 / single                      // 单核操作速率
	tn := float64(p*1000) / duration.Seconds() // 多核操作速率
	efficiency := tn / t1 / float64(proc)      // 多核效率
	return BenchmarkResult{
		Name:       sb.Name(),
		Category:   sb.Category(),
		Score:      tn / 10.0,
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
	}
}

func stringTest(operations int) {
	words := make([]string, 100)
	for i := 0; i < 100; i++ {
		words[i] = "benchmark_string_processing_performance_test_data"
	}
	for i := 0; i < operations; i++ {
		// 字符串拼接
		result := ""
		for _, word := range words {
			result += word
		}
		// 字符串搜索
		_ = strings.Contains(result, "performance")
		_ = strings.Index(result, "test")

		// 字符串替换
		_ = strings.ReplaceAll(result, "benchmark", "performance")
	}
}
