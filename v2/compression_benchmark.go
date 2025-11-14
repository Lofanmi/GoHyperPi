package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io"
	"strings"
	"sync"
	"time"
)

// CompressionBenchmark 压缩解压缩性能测试
type CompressionBenchmark struct{}

// NewCompressionBenchmark 创建压缩测试实例
func NewCompressionBenchmark() *CompressionBenchmark {
	return &CompressionBenchmark{}
}

func (cb *CompressionBenchmark) Name() string {
	return "压缩性能测试（Compression）"
}

func (cb *CompressionBenchmark) Description() string {
	return "测试压缩和解压缩性能"
}

func (cb *CompressionBenchmark) Category() string {
	return "压缩性能"
}

func (cb *CompressionBenchmark) Run(proc, times int) BenchmarkResult {
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)
	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			compressionTest(100) // 100次压缩操作
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
	t1 := 100.0 / single                      // 单核压缩速率（ops/s）
	tn := float64(p*100) / duration.Seconds() // 多核压缩速率（ops/s）
	efficiency := tn / t1 / float64(proc)     // 多核效率
	return BenchmarkResult{
		Name:       cb.Name(),
		Category:   cb.Category(),
		Score:      tn / 10.0, // 归一化得分
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
	}
}

func compressionTest(operations int) {
	// 创建测试数据
	text := strings.Repeat("The quick brown fox jumps over the lazy dog. ", 1000)
	data := []byte(text)
	for i := 0; i < operations/2; i++ {
		// Gzip压缩测试
		var buf bytes.Buffer
		gw := gzip.NewWriter(&buf)
		gw.Write(data)
		gw.Close()
		// Gzip解压缩测试
		gr, _ := gzip.NewReader(&buf)
		_, _ = io.ReadAll(gr)
		gr.Close()
	}
	for i := 0; i < operations/2; i++ {
		// Zlib压缩测试
		var buf bytes.Buffer
		zw := zlib.NewWriter(&buf)
		zw.Write(data)
		zw.Close()

		// Zlib解压缩测试
		zr, _ := zlib.NewReader(&buf)
		_, _ = io.ReadAll(zr)
		zr.Close()
	}
}

// GzipBenchmark Gzip专用测试
type GzipBenchmark struct{}

func NewGzipBenchmark() *GzipBenchmark {
	return &GzipBenchmark{}
}

func (gb *GzipBenchmark) Name() string {
	return "Gzip压缩（Gzip）"
}

func (gb *GzipBenchmark) Description() string {
	return "测试Gzip压缩算法性能"
}

func (gb *GzipBenchmark) Category() string {
	return "压缩性能"
}

func (gb *GzipBenchmark) Run(proc, times int) BenchmarkResult {
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)
	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			gzipTest(50)
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
	t1 := 50.0 / single                      // 单核操作速率
	tn := float64(p*50) / duration.Seconds() // 多核操作速率
	efficiency := tn / t1 / float64(proc)    // 多核效率
	return BenchmarkResult{
		Name:       gb.Name(),
		Category:   gb.Category(),
		Score:      tn / 10.0,
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
	}
}

func gzipTest(operations int) {
	text := strings.Repeat("Performance testing data compression algorithms with various input sizes and patterns. ", 2000)
	data := []byte(text)
	for i := 0; i < operations; i++ {
		var buf bytes.Buffer
		gw := gzip.NewWriter(&buf)
		_, _ = gw.Write(data)
		gw.Close()
		gr, _ := gzip.NewReader(&buf)
		_, _ = io.ReadAll(gr)
		gr.Close()
	}
}
