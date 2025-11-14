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
	Name       string
	Category   string
	Score      float64       // 综合得分
	Duration   time.Duration // 总耗时
	SingleRate float64       // 单核性能指标
	MultiRate  float64       // 多核性能指标
	Efficiency float64       // 多核效率 (MultiRate/SingleRate/Proc)
	Proc       int           // 使用的核心数
	Times      int           // 运行次数
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
			NewGzipBenchmark(),             // Gzip专用测试
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
		// 显示当前测试进度
		fmt.Printf("正在测试 %s ...\n", benchmark.Name())

		// 自适应测试：确保每个测试至少运行5秒
		result := bs.runAdaptiveBenchmark(benchmark, proc, times)
		fmt.Printf("测试完成（用时 %s）\n", formatDuration(result.Duration.Seconds()))

		results = append(results, result)
	}
	fmt.Println()
	return results
}

// runAdaptiveBenchmark 自适应运行基准测试，确保稳定性和足够长的测试时间
func (bs *BenchmarkSuite) runAdaptiveBenchmark(benchmark Benchmark, proc, times int) BenchmarkResult {
	const targetDuration = 5 * time.Second // 目标测试时间5秒
	const minRounds = 10                   // 最少测试轮次
	const maxRounds = 100                  // 最多测试轮次

	var results []BenchmarkResult
	var totalDuration time.Duration

	// 先进行一轮测试，估算所需时间
	sampleResult := benchmark.Run(proc, 1)
	estimatedTimePerRound := sampleResult.Duration

	// 计算需要的轮次
	estimatedRounds := int(targetDuration / estimatedTimePerRound)
	if estimatedRounds < minRounds {
		estimatedRounds = minRounds
	}
	if estimatedRounds > maxRounds {
		estimatedRounds = maxRounds
	}

	// 运行多轮测试
	for i := 0; i < estimatedRounds; i++ {
		result := benchmark.Run(proc, times)
		results = append(results, result)
		totalDuration += result.Duration

		// 如果已经运行足够时间，可以提前结束
		if totalDuration >= targetDuration && i >= minRounds-1 {
			break
		}
	}

	// 计算平均值，去掉最高和最低的20%结果（提高稳定性）
	stableResults := bs.filterOutliers(results)

	return bs.calculateAverageResult(benchmark, stableResults, proc, times)
}

// filterOutliers 过滤异常值，去掉最高和最低的20%
func (bs *BenchmarkSuite) filterOutliers(results []BenchmarkResult) []BenchmarkResult {
	if len(results) <= 3 {
		return results
	}

	// 按分数排序
	sorted := make([]BenchmarkResult, len(results))
	copy(sorted, results)

	// 简单排序（按分数）
	for i := 0; i < len(sorted)-1; i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[i].Score > sorted[j].Score {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}

	// 去掉最低的20%和最高的20%
	removeCount := len(sorted) / 5
	if removeCount == 0 {
		removeCount = 1
	}

	start := removeCount
	end := len(sorted) - removeCount
	if start >= end {
		start = 0
		end = len(sorted)
	}

	return sorted[start:end]
}

// calculateAverageResult 计算平均结果
func (bs *BenchmarkSuite) calculateAverageResult(benchmark Benchmark, results []BenchmarkResult, proc, times int) BenchmarkResult {
	var totalScore, totalSingle, totalMulti, totalEfficiency float64
	var totalDuration time.Duration

	for _, result := range results {
		totalScore += result.Score
		totalSingle += result.SingleRate
		totalMulti += result.MultiRate
		totalEfficiency += result.Efficiency
		totalDuration += result.Duration
	}

	count := float64(len(results))
	return BenchmarkResult{
		Name:       benchmark.Name(),
		Category:   benchmark.Category(),
		Score:      totalScore / count,
		SingleRate: totalSingle / count,
		MultiRate:  totalMulti / count,
		Efficiency: totalEfficiency / count,
		Duration:   totalDuration / time.Duration(len(results)),
		Proc:       proc,
		Times:      times,
	}
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
