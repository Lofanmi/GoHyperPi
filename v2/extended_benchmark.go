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
)

// TrigBenchmark 三角函数测试
type TrigBenchmark struct {
	*BaseBenchmark
}

func NewTrigBenchmark() *TrigBenchmark {
	return &TrigBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"三角函数计算（Trigonometric Functions）",
			"测试三角函数性能",
			"浮点性能",
			trigTest,
			5000000, // 500万次三角函数运算
		),
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
type BitOperationsBenchmark struct {
	*BaseBenchmark
}

func NewBitOperationsBenchmark() *BitOperationsBenchmark {
	return &BitOperationsBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"位运算测试（Bit Operations）",
			"测试位运算性能",
			"计算密集型",
			bitTest,
			1000000000, // 10亿次操作
		),
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
type AdvancedCryptoBenchmark struct {
	*BaseBenchmark
}

func NewAdvancedCryptoBenchmark() *AdvancedCryptoBenchmark {
	return &AdvancedCryptoBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"高级加密算法（Advanced Cryptography）",
			"测试高级加密算法性能",
			"加密性能",
			advancedCryptoTest,
			100000, // 10万次操作
		),
	}
}

func advancedCryptoTest(operations int) {
	data := make([]byte, 1024)
	key := make([]byte, 32)
	iv := make([]byte, aes.BlockSize)
	_, _ = rand.Read(key)
	_, _ = rand.Read(iv)
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
type IntegerBenchmark struct {
	*BaseBenchmark
}

func NewIntegerBenchmark() *IntegerBenchmark {
	return &IntegerBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"整数运算测试（Integer Operations）",
			"测试整数运算性能",
			"计算密集型",
			integerTest,
			1000000000, // 10亿次操作
		),
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
type BinaryBenchmark struct {
	*BaseBenchmark
}

func NewBinaryBenchmark() *BinaryBenchmark {
	return &BinaryBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"二进制处理（Binary Processing）",
			"测试二进制数据处理性能",
			"算法性能",
			binaryTest,
			1000000000, // 10亿次操作
		),
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
