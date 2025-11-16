package main

import (
	"math/rand"
	"sort"
	"strings"
)

// SortingBenchmark 排序算法性能测试
type SortingBenchmark struct {
	*BaseBenchmark
}

// NewSortingBenchmark 创建排序测试实例
func NewSortingBenchmark() *SortingBenchmark {
	testFunc := func(workload int) {
		sortingTest(workload)
	}

	return &SortingBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"排序算法测试（Sorting Algorithms）",
			"测试各种排序算法性能",
			"算法性能",
			testFunc,
			1000000, // 排序100000个元素
		),
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
type StringBenchmark struct {
	*BaseBenchmark
}

func NewStringBenchmark() *StringBenchmark {
	testFunc := func(workload int) {
		stringTest(workload)
	}

	return &StringBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"字符串处理（String Processing）",
			"测试字符串操作性能",
			"算法性能",
			testFunc,
			1000, // 1000次字符串操作
		),
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
