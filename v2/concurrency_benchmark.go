package main

import (
	"sync"
	"time"
)

// ConcurrencyBenchmark 并发处理性能测试
type ConcurrencyBenchmark struct{}

// NewConcurrencyBenchmark 创建并发测试实例
func NewConcurrencyBenchmark() *ConcurrencyBenchmark {
	return &ConcurrencyBenchmark{}
}

func (cb *ConcurrencyBenchmark) Name() string {
	return "并发测试（Concurrency Test）"
}

func (cb *ConcurrencyBenchmark) Description() string {
	return "测试并发处理和同步性能"
}

func (cb *ConcurrencyBenchmark) Category() string {
	return "并发性能"
}

func (cb *ConcurrencyBenchmark) Run(proc, times int) BenchmarkResult {
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)
	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			concurrencyTest(1000000) // 100万次操作
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
	t1 := 1000000.0 / single                     // 单核操作速率（ops/s）
	tn := float64(p*1000000) / duration.Seconds() // 多核操作速率（ops/s）

	// 避免除零错误
	var efficiency float64
	if t1 > 0 {
		efficiency = tn / t1 / float64(proc)      // 多核效率
	} else {
		efficiency = 0.0
	}
	return BenchmarkResult{
		Name:       cb.Name(),
		Category:   cb.Category(),
		Score:      tn / 1000.0, // 归一化得分（kops/s）
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
	}
}

func concurrencyTest(operations int) {
	var wg sync.WaitGroup
	counter := 0
	mu := sync.Mutex{}
	concurrentGoroutines := 10
	// 并发计数器测试
	for i := 0; i < operations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
		if i%concurrentGoroutines == 0 {
			wg.Wait()
		}
	}
	wg.Wait()
	// 通道通信测试
	ch := make(chan int, 100)
	done := make(chan bool)
	go func() {
		for val := range ch {
			_ = val
		}
		done <- true
	}()
	for i := 0; i < 10000; i++ {
		ch <- i
	}
	close(ch)
	<-done
	// WaitGroup测试
	for i := 0; i < 100; i++ {
		wg.Add(10)
		for j := 0; j < 10; j++ {
			go func() {
				defer wg.Done()
				time.Sleep(time.Microsecond)
			}()
		}
		wg.Wait()
	}
	_ = counter
}

// ChannelBenchmark 通道通信测试
type ChannelBenchmark struct{}

func NewChannelBenchmark() *ChannelBenchmark {
	return &ChannelBenchmark{}
}

func (ccb *ChannelBenchmark) Name() string {
	return "通道通信测试（Channel Communication）"
}

func (ccb *ChannelBenchmark) Description() string {
	return "测试Goroutine间通信性能"
}

func (ccb *ChannelBenchmark) Category() string {
	return "并发性能"
}

func (ccb *ChannelBenchmark) Run(proc, times int) BenchmarkResult {
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)
	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			channelTest(100000) // 10万次消息
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
	t1 := 100000.0 / single                     // 单核消息速率（msg/s）
	tn := float64(p*100000) / duration.Seconds() // 多核消息速率（msg/s）

	// 避免除零错误
	var efficiency float64
	if t1 > 0 {
		efficiency = tn / t1 / float64(proc)      // 多核效率
	} else {
		efficiency = 0.0
	}
	return BenchmarkResult{
		Name:       ccb.Name(),
		Category:   ccb.Category(),
		Score:      tn / 10000.0, // 归一化得分
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
	}
}

func channelTest(messages int) {
	ch := make(chan int, 100)
	var wg sync.WaitGroup
	// 启动接收者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range ch {
		}
	}()
	// 发送消息
	for i := 0; i < messages; i++ {
		ch <- i
	}
	close(ch)

	wg.Wait()
}