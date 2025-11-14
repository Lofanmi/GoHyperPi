package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"sync"
	"time"
)

// CryptoBenchmark 加密运算性能测试
type CryptoBenchmark struct{}

// NewCryptoBenchmark 创建加密测试实例
func NewCryptoBenchmark() *CryptoBenchmark {
	return &CryptoBenchmark{}
}

func (cb *CryptoBenchmark) Name() string {
	return "加密算法测试（Cryptography）"
}

func (cb *CryptoBenchmark) Description() string {
	return "测试哈希和加密算法性能"
}

func (cb *CryptoBenchmark) Category() string {
	return "加密性能"
}

func (cb *CryptoBenchmark) Run(proc, times int) BenchmarkResult {
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)
	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			cryptoTest(50000) // 5万次操作
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
	t1 := 50000.0 / single                    // 单核操作速率（ops/s）
	tn := float64(p*50000) / duration.Seconds() // 多核操作速率（ops/s）

	// 避免除零错误
	var efficiency float64
	if t1 > 0 {
		efficiency = tn / t1 / float64(proc)   // 多核效率
	} else {
		efficiency = 0.0
	}
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

func cryptoTest(operations int) {
	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte(i)
	}
	// MD5哈希测试
	for i := 0; i < operations/4; i++ {
		hash := md5.Sum(data)
		_ = hex.EncodeToString(hash[:])
	}
	// SHA256哈希测试
	for i := 0; i < operations/4; i++ {
		hash := sha256.Sum256(data)
		_ = hex.EncodeToString(hash[:])
	}
	// SHA1哈希测试
	for i := 0; i < operations/4; i++ {
		hash := sha1.Sum(data)
		_ = hex.EncodeToString(hash[:])
	}
	// AES加密测试
	block, _ := aes.NewCipher([]byte("1234567890123456"))
	stream := cipher.NewCTR(block, make([]byte, aes.BlockSize))
	encrypted := make([]byte, len(data))
	for i := 0; i < operations/4; i++ {
		stream.XORKeyStream(encrypted, data)
	}
}

// HashBenchmark 哈希运算专用测试
type HashBenchmark struct{}

func NewHashBenchmark() *HashBenchmark {
	return &HashBenchmark{}
}

func (hb *HashBenchmark) Name() string {
	return "哈希函数测试（Hash Functions）"
}

func (hb *HashBenchmark) Description() string {
	return "测试各类哈希函数性能"
}

func (hb *HashBenchmark) Category() string {
	return "加密性能"
}

func (hb *HashBenchmark) Run(proc, times int) BenchmarkResult {
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)
	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			hashTest(200000) // 20万次哈希
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
	t1 := 200000.0 / single                    // 单核哈希速率（hash/s）
	tn := float64(p*200000) / duration.Seconds() // 多核哈希速率（hash/s）

	// 避免除零错误
	var efficiency float64
	if t1 > 0 {
		efficiency = tn / t1 / float64(proc)   // 多核效率
	} else {
		efficiency = 0.0
	}
	return BenchmarkResult{
		Name:       hb.Name(),
		Category:   hb.Category(),
		Score:      tn / 1000.0, // 归一化得分
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
	}
}

func hashTest(operations int) {
	data := make([]byte, 1024)
	rand.Read(data)
	for i := 0; i < operations/4; i++ {
		_ = md5.Sum(data)
	}
	for i := 0; i < operations/4; i++ {
		_ = sha1.Sum(data)
	}
	for i := 0; i < operations/4; i++ {
		_ = sha256.Sum256(data)
	}
	for i := 0; i < operations/4; i++ {
		_ = sha256.Sum256(data)
	}
}