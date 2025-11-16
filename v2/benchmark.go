package main

import (
	"fmt"
	"time"
)

// Benchmark 测试项目接口
type Benchmark interface {
	Name() string
	Description() string
	Category() string
	Run(proc, times int) BenchmarkResult
}

// BenchmarkResult 测试结果
type BenchmarkResult struct {
	Name           string
	Category       string
	Duration       time.Duration
	SingleDuration time.Duration // 单核性能指标
	MultiDuration  time.Duration // 多核性能指标
	Ratio          float64       // 倍率
	Score          float64       // 综合得分
	Proc           int           // 使用的核心数
	Times          int           // 运行次数
}

// BenchmarkSuite 测试套件
type BenchmarkSuite struct {
	benchmarks []Benchmark
}

// NewBenchmarkSuite 创建新的测试套件
func NewBenchmarkSuite() *BenchmarkSuite {
	return &BenchmarkSuite{
		benchmarks: []Benchmark{
			NewComputeBenchmark(),          // 计算密集型测试（Pi计算）
			NewBitOperationsBenchmark(),    // 位运算测试
			NewIntegerBenchmark(),          // 整数运算测试
			NewMemoryBenchmark(),           // 内存访问测试
			NewMemorySequentialBenchmark(), // 顺序内存访问测试
			NewConcurrencyBenchmark(),      // 并发处理测试
			NewChannelBenchmark(),          // 通道通信测试
			NewCryptoBenchmark(),           // 加密运算测试
			NewAdvancedCryptoBenchmark(),   // 高级加密测试
			NewHashBenchmark(),             // 哈希运算测试
			NewFloatBenchmark(),            // 浮点运算测试
			NewTrigBenchmark(),             // 三角函数测试
			NewMatrixBenchmark(),           // 矩阵运算测试
			NewCompressionBenchmark(),      // 压缩性能测试
			NewSortingBenchmark(),          // 排序算法测试
			NewStringBenchmark(),           // 字符串处理测试
			NewBinaryBenchmark(),           // 二进制处理测试
		},
	}
}

// AddBenchmark 添加测试项目
func (bs *BenchmarkSuite) AddBenchmark(benchmark Benchmark) {
	bs.benchmarks = append(bs.benchmarks, benchmark)
}

// RunBenchmarks 运行所有测试
func (bs *BenchmarkSuite) RunBenchmarks(proc, times int) []BenchmarkResult {
	var results []BenchmarkResult
	fmt.Println()
	for _, benchmark := range bs.benchmarks {
		fmt.Printf("正在测试 %s ...\n", benchmark.Name())
		result := benchmark.Run(proc, times)
		fmt.Printf("测试完成（用时 %s）\n", formatDuration(result.Duration.Seconds()))
		results = append(results, result)
	}
	fmt.Println()
	return results
}

// GetCategories 获取所有测试类别
func (bs *BenchmarkSuite) GetCategories() []string {
	categories := make(map[string]bool)
	for _, benchmark := range bs.benchmarks {
		categories[benchmark.Category()] = true
	}
	var result []string
	for category := range categories {
		result = append(result, category)
	}
	return result
}
