package main

import (
	"math/rand"
	"time"
)

// MemoryBenchmark 内存访问性能测试
type MemoryBenchmark struct {
	*BaseBenchmark
}

// NewMemoryBenchmark 创建内存性能测试实例
func NewMemoryBenchmark() *MemoryBenchmark {
	testFunc := func(workload int) {
		// 将workload映射为内存大小（MB）
		memorySize := int64(workload) * 1024 * 1024 // workload MB
		memoryTest(memorySize)
	}

	return &MemoryBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"内存访问测试（Memory Access）",
			"测试内存读写性能和缓存效率",
			"内存性能",
			testFunc,
			100, // 100MB默认大小
		),
	}
}

// memoryTest 内存测试函数
func memoryTest(size int64) {
	data := make([]byte, size)
	for i := int64(0); i < size; i++ {
		data[i] = byte(i % 256)
	}
	sum := 0
	for i := 0; i < len(data); i++ {
		sum += int(data[i])
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(data)/100; i++ {
		idx := r.Intn(len(data))
		sum += int(data[idx])
	}
	_ = sum
}

// MemorySequentialBenchmark 顺序内存访问测试
type MemorySequentialBenchmark struct {
	*BaseBenchmark
}

func NewMemorySequentialBenchmark() *MemorySequentialBenchmark {
	testFunc := func(workload int) {
		// 将workload映射为内存大小（MB）
		memorySize := int64(workload) * 1024 * 1024 // workload MB
		sequentialMemoryTest(memorySize)
	}

	return &MemorySequentialBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"顺序内存访问（Sequential Memory）",
			"测试顺序内存访问性能",
			"内存性能",
			testFunc,
			100, // 100MB默认大小
		),
	}
}

func sequentialMemoryTest(size int64) {
	data := make([]byte, size)
	for i := int64(0); i < size; i++ {
		data[i] = byte(i)
	}
	sum := uint64(0)
	for i := 0; i < len(data); i++ {
		sum += uint64(data[i])
	}
	_ = sum
}
