package main

import (
	"sync"
	"time"
)

// ConcurrencyBenchmark 并发处理性能测试
type ConcurrencyBenchmark struct {
	*BaseBenchmark
}

// NewConcurrencyBenchmark 创建并发测试实例
func NewConcurrencyBenchmark() *ConcurrencyBenchmark {
	return &ConcurrencyBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"并发测试（Concurrency Test）",
			"测试并发处理和同步性能",
			"并发性能",
			concurrencyTest,
			1000000, // 100万次操作
		),
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
type ChannelBenchmark struct {
	*BaseBenchmark
}

func NewChannelBenchmark() *ChannelBenchmark {
	return &ChannelBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"通道通信测试（Channel Communication）",
			"测试Goroutine间通信性能",
			"并发性能",
			channelTest,
			1000000, // 100万次消息
		),
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
