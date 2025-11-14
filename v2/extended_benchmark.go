package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/binary"
	"math"
	"math/bits"
	"sync"
	"time"
)

// TrigBenchmark 三角函数测试
type TrigBenchmark struct{}

func NewTrigBenchmark() *TrigBenchmark {
	return &TrigBenchmark{}
}

func (tb *TrigBenchmark) Name() string {
	return "三角函数计算（Trigonometric Functions）"
}

func (tb *TrigBenchmark) Description() string {
	return "测试三角函数性能"
}

func (tb *TrigBenchmark) Category() string {
	return "浮点性能"
}

func (tb *TrigBenchmark) Run(proc, times int) BenchmarkResult {
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)

	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			trigTest(2000000) // 200万次三角函数运算
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

	t1 := 2000000.0 / single                      // 单核操作速率
	tn := float64(p*2000000) / duration.Seconds() // 多核操作速率

	// 避免除零错误
	var efficiency float64
	if t1 > 0 {
		efficiency = tn / t1 / float64(proc)      // 多核效率
	} else {
		efficiency = 0.0
	}

	return BenchmarkResult{
		Name:       tb.Name(),
		Category:   tb.Category(),
		Score:      tn / 1000.0,
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
	}
}

func trigTest(operations int) {
	for i := 0; i < operations; i++ {
		x := float64(i) * 0.001
		_ = math.Sin(x)
		_ = math.Cos(x)
		_ = math.Tan(x)
		_ = math.Asin(math.Sin(x))
		_ = math.Acos(math.Cos(x))
		_ = math.Atan(math.Tan(x))
	}
}

// BitOperationsBenchmark 位运算测试
type BitOperationsBenchmark struct{}

func NewBitOperationsBenchmark() *BitOperationsBenchmark {
	return &BitOperationsBenchmark{}
}

func (bob *BitOperationsBenchmark) Name() string {
	return "位运算测试（Bit Operations）"
}

func (bob *BitOperationsBenchmark) Description() string {
	return "测试位运算性能"
}

func (bob *BitOperationsBenchmark) Category() string {
	return "计算密集型"
}

func (bob *BitOperationsBenchmark) Run(proc, times int) BenchmarkResult {
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)

	// 增加运算量，确保测试时间足够长
	operations := 50000000 // 5000万次操作
	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			bitTest(operations)
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

	t1 := float64(operations) / single               // 单核操作速率
	tn := float64(p*operations) / duration.Seconds() // 多核操作速率

	// 避免除零错误
	var efficiency float64
	if t1 > 0 {
		efficiency = tn / t1 / float64(proc)         // 多核效率
	} else {
		efficiency = 0.0
	}

	return BenchmarkResult{
		Name:       bob.Name(),
		Category:   bob.Category(),
		Score:      tn / 1000.0,
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
	}
}

func bitTest(operations int) {
	for i := 0; i < operations; i++ {
		n := uint32(i)
		_ = bits.LeadingZeros32(n)
		_ = bits.TrailingZeros32(n)
		_ = bits.OnesCount32(n)
		_ = bits.ReverseBytes32(n)
		_ = bits.RotateLeft32(n, 5)
		_ = n & (n - 1)
		_ = n | (n + 1)
		_ = n ^ (n + 1)
	}
}

// AdvancedCryptoBenchmark 高级加密测试
type AdvancedCryptoBenchmark struct{}

func NewAdvancedCryptoBenchmark() *AdvancedCryptoBenchmark {
	return &AdvancedCryptoBenchmark{}
}

func (acb *AdvancedCryptoBenchmark) Name() string {
	return "高级加密算法（Advanced Cryptography）"
}

func (acb *AdvancedCryptoBenchmark) Description() string {
	return "测试高级加密算法性能"
}

func (acb *AdvancedCryptoBenchmark) Category() string {
	return "加密性能"
}

func (acb *AdvancedCryptoBenchmark) Run(proc, times int) BenchmarkResult {
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)

	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			advancedCryptoTest(500)
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

	t1 := 500.0 / single                      // 单核操作速率
	tn := float64(p*500) / duration.Seconds() // 多核操作速率

	// 避免除零错误
	var efficiency float64
	if t1 > 0 {
		efficiency = tn / t1 / float64(proc)   // 多核效率
	} else {
		efficiency = 0.0
	}

	return BenchmarkResult{
		Name:       acb.Name(),
		Category:   acb.Category(),
		Score:      tn / 10.0,
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
	}
}

func advancedCryptoTest(operations int) {
	data := make([]byte, 1024)
	key := make([]byte, 32)
	iv := make([]byte, aes.BlockSize)
	rand.Read(key)
	rand.Read(iv)
	for i := 0; i < operations/2; i++ {
		hash := sha512.Sum512(data)
		_ = hash
		mac := hmac.New(sha512.New, key)
		mac.Write(data)
		_ = mac.Sum(nil)
	}
	for i := 0; i < operations/2; i++ {
		block, _ := aes.NewCipher(key)
		aesgcm, _ := cipher.NewGCM(block)
		nonce := make([]byte, aesgcm.NonceSize())
		ciphertext := aesgcm.Seal(nil, nonce, data, nil)
		_ = ciphertext
		cbc := cipher.NewCBCEncrypter(block, iv)
		encrypted := make([]byte, len(data))
		cbc.CryptBlocks(encrypted, data)
	}
}

// IntegerBenchmark 整数运算测试
type IntegerBenchmark struct{}

func NewIntegerBenchmark() *IntegerBenchmark {
	return &IntegerBenchmark{}
}

func (ib *IntegerBenchmark) Name() string {
	return "整数运算测试（Integer Operations）"
}

func (ib *IntegerBenchmark) Description() string {
	return "测试整数运算性能"
}

func (ib *IntegerBenchmark) Category() string {
	return "计算密集型"
}

func (ib *IntegerBenchmark) Run(proc, times int) BenchmarkResult {
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)

	// 增加运算量，确保测试时间足够长
	operations := 100000000 // 1亿次操作
	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			integerTest(operations)
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

	t1 := float64(operations) / single               // 单核操作速率
	tn := float64(p*operations) / duration.Seconds() // 多核操作速率

	// 避免除零错误
	var efficiency float64
	if t1 > 0 {
		efficiency = tn / t1 / float64(proc)         // 多核效率
	} else {
		efficiency = 0.0
	}

	return BenchmarkResult{
		Name:       ib.Name(),
		Category:   ib.Category(),
		Score:      tn / 1000.0,
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
	}
}

func integerTest(operations int) {
	for i := 0; i < operations; i++ {
		a := int64(i)
		b := int64(i + 1)
		c := int64(i + 2)
		_ = a + b*c
		_ = (a + b) / (c + 1)
		_ = a ^ b ^ c
		_ = (a + b) & (c | a)
		_ = a*b + c*a
		_ = (a << 3) >> 2
		_ = a % (b + 1)
	}
}

// BinaryBenchmark 二进制数据处理测试
type BinaryBenchmark struct{}

func NewBinaryBenchmark() *BinaryBenchmark {
	return &BinaryBenchmark{}
}

func (bb *BinaryBenchmark) Name() string {
	return "二进制处理（Binary Processing）"
}

func (bb *BinaryBenchmark) Description() string {
	return "测试二进制数据处理性能"
}

func (bb *BinaryBenchmark) Category() string {
	return "算法性能"
}

func (bb *BinaryBenchmark) Run(proc, times int) BenchmarkResult {
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)

	// 增加运算量，确保测试时间足够长
	operations := 50000000 // 5000万次操作
	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			binaryTest(operations)
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

	t1 := float64(operations) / single               // 单核操作速率
	tn := float64(p*operations) / duration.Seconds() // 多核操作速率

	// 避免除零错误
	var efficiency float64
	if t1 > 0 {
		efficiency = tn / t1 / float64(proc)         // 多核效率
	} else {
		efficiency = 0.0
	}

	return BenchmarkResult{
		Name:       bb.Name(),
		Category:   bb.Category(),
		Score:      tn / 10.0,
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
	}
}

func binaryTest(operations int) {
	for i := 0; i < operations; i++ {
		var buf [8]byte
		binary.BigEndian.PutUint64(buf[:], uint64(i))
		_ = binary.BigEndian.Uint64(buf[:])
		binary.LittleEndian.PutUint32(buf[:4], uint32(i))
		_ = binary.LittleEndian.Uint32(buf[:4])
		num := uint64(i)
		_ = (num & 0xFF00) >> 8
		_ = (num << 16) & 0xFFFF0000
		_ = ^num
	}
}
