package main

import (
	"flag"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/klauspost/cpuid/v2"
)

func main() {
	var (
		proc     int
		times    int
		verbose  bool
		cpuInfo  bool
		category string
		output   string
	)
	P := runtime.GOMAXPROCS(0)
	flag.IntVar(&proc, "proc", P, "Processor count")
	flag.IntVar(&times, "times", 1, "Test times per processor")
	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.BoolVar(&cpuInfo, "cpu", true, "Show CPU information")
	flag.StringVar(&category, "category", "", "Run specific category only")
	flag.StringVar(&output, "output", "", "Output report to file")
	flag.Parse()
	// 显示CPU信息
	if cpuInfo {
		printCPUInfo()
	}
	// 创建测试套件
	suite := NewBenchmarkSuite()
	calculator := NewScoreCalculator()
	// 过滤特定类别
	if category != "" {
		filteredSuite := &BenchmarkSuite{}
		for _, benchmark := range suite.benchmarks {
			if benchmark.Category() == category {
				filteredSuite.AddBenchmark(benchmark)
			}
		}
		suite = filteredSuite
	}
	fmt.Println("开始运行性能测试...")
	startTime := time.Now()
	// 运行基准测试
	results := suite.RunBenchmarks(proc, times)
	totalDuration := time.Since(startTime)
	if verbose {
		fmt.Println("详细测试结果:")
		for _, result := range results {
			fmt.Printf("  %-6s | %-32s | 得分: %8.2f | 单核: %8.2f | 多核: %8.2f | 多核/单核: %.2f | 耗时: %v\n",
				result.Category, result.Name, result.Score, result.SingleRate, result.MultiRate, result.MultiRate/result.SingleRate, result.Duration)
		}
		fmt.Println()
	}
	// 生成并显示报告
	report := calculator.GenerateReport(results)
	fmt.Println(report)
	// 显示总耗时
	fmt.Printf("总测试时间: %v\n", totalDuration)
	// 输出到文件
	if output != "" {
		err := writeReportToFile(report, output)
		if err != nil {
			fmt.Printf("保存报告失败: %v\n", err)
		} else {
			fmt.Printf("报告已保存到: %s\n", output)
		}
	}
}

func printCPUInfo() {
	fmt.Println("=== CPU 信息 ===")
	fmt.Printf("CPU 名称: %s\n", cpuid.CPU.BrandName)
	fmt.Printf("物理核心: %d\n", cpuid.CPU.PhysicalCores)
	fmt.Printf("每核心线程数: %d\n", cpuid.CPU.ThreadsPerCore)
	fmt.Printf("逻辑核心: %d\n", cpuid.CPU.LogicalCores)
	fmt.Printf("CPU 系列: %d, 型号: %d, 厂商ID: %s\n", cpuid.CPU.Family, cpuid.CPU.Model, cpuid.CPU.VendorID)

	features := strings.Join(cpuid.CPU.FeatureSet(), ", ")
	fmt.Printf("CPU 指令集: %s\n", features)
	fmt.Printf("缓存行大小: %d 字节\n", cpuid.CPU.CacheLine)
	fmt.Printf("L1 数据缓存: %d 字节\n", cpuid.CPU.Cache.L1D)
	fmt.Printf("L1 指令缓存: %d 字节\n", cpuid.CPU.Cache.L1I)
	fmt.Printf("L2 缓存: %d 字节\n", cpuid.CPU.Cache.L2)
	if cpuid.CPU.Cache.L3 > 0 {
		fmt.Printf("L3 缓存: %d 字节\n", cpuid.CPU.Cache.L3)
	}
	fmt.Printf("CPU 频率: %d Hz\n", cpuid.CPU.Hz)
	fmt.Printf("操作系统: %s %s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Println()
}
