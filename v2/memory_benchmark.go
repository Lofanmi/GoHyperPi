package main

import (
	"math/rand"
	"sync"
	"time"
)

// MemoryBenchmark 内存访问性能测试
type MemoryBenchmark struct {
	size int64 // 测试数据大小（字节）
}

// NewMemoryBenchmark 创建内存性能测试实例
func NewMemoryBenchmark() *MemoryBenchmark {
	return &MemoryBenchmark{
		size: 100 * 1024 * 1024, // 默认100MB
	}
}

func (mb *MemoryBenchmark) Name() string {
	return "内存访问测试（Memory Access）"
}

func (mb *MemoryBenchmark) Description() string {
	return "测试内存读写性能和缓存效率"
}

func (mb *MemoryBenchmark) Category() string {
	return "内存性能"
}

func (mb *MemoryBenchmark) Run(proc, times int) BenchmarkResult {
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)

	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			memoryTest(mb.size)
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

	// 计算内存访问速率（GB/s）
	t1 := float64(mb.size) / (1024*1024*1024) / single                           // 单核速率
	tn := float64(int64(p)*mb.size) / (1024*1024*1024) / duration.Seconds() // 多核速率
	efficiency := tn / t1 / float64(proc)                         // 多核效率

	return BenchmarkResult{
		Name:       mb.Name(),
		Category:   mb.Category(),
		Score:      tn * 10, // 归一化得分（GB/s * 10）
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
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
type MemorySequentialBenchmark struct{}

func NewMemorySequentialBenchmark() *MemorySequentialBenchmark {
	return &MemorySequentialBenchmark{}
}

func (msb *MemorySequentialBenchmark) Name() string {
	return "顺序内存访问（Sequential Memory）"
}

func (msb *MemorySequentialBenchmark) Description() string {
	return "测试顺序内存访问性能"
}

func (msb *MemorySequentialBenchmark) Category() string {
	return "内存性能"
}

func (msb *MemorySequentialBenchmark) Run(proc, times int) BenchmarkResult {
	size := int64(50 * 1024 * 1024) // 50MB
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)

	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			sequentialMemoryTest(size)
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

	t1 := float64(size) / (1024*1024*1024) / single
	tn := float64(int64(p)*size) / (1024*1024*1024) / duration.Seconds()
	efficiency := tn / t1 / float64(proc)

	return BenchmarkResult{
		Name:       msb.Name(),
		Category:   msb.Category(),
		Score:      tn * 10,
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
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