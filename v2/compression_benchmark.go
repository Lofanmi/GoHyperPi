package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io"
	"strings"
)

// CompressionBenchmark 压缩解压缩性能测试
type CompressionBenchmark struct {
	*BaseBenchmark
}

// NewCompressionBenchmark 创建压缩测试实例
func NewCompressionBenchmark() *CompressionBenchmark {
	testFunc := func(workload int) {
		compressionTest(workload)
	}

	return &CompressionBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"压缩性能测试（Compression）",
			"测试压缩和解压缩性能",
			"压缩性能",
			testFunc,
			500, // 500次压缩操作
		),
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
		_, _ = gw.Write(data)
		_ = gw.Close()
		// Gzip解压缩测试
		gr, _ := gzip.NewReader(&buf)
		_, _ = io.ReadAll(gr)
		_ = gr.Close()
	}
	for i := 0; i < operations/2; i++ {
		// Zlib压缩测试
		var buf bytes.Buffer
		zw := zlib.NewWriter(&buf)
		_, _ = zw.Write(data)
		_ = zw.Close()

		// Zlib解压缩测试
		zr, _ := zlib.NewReader(&buf)
		_, _ = io.ReadAll(zr)
		_ = zr.Close()
	}
}
