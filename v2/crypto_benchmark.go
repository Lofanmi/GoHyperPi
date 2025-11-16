package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// CryptoBenchmark 加密运算性能测试
type CryptoBenchmark struct {
	*BaseBenchmark
}

// NewCryptoBenchmark 创建加密测试实例
func NewCryptoBenchmark() *CryptoBenchmark {
	testFunc := func(workload int) {
		cryptoTest(workload)
	}

	return &CryptoBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"加密算法测试（Cryptography）",
			"测试哈希和加密算法性能",
			"加密性能",
			testFunc,
			200000, // 20万次操作
		),
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
type HashBenchmark struct {
	*BaseBenchmark
}

func NewHashBenchmark() *HashBenchmark {
	testFunc := func(workload int) {
		hashTest(workload)
	}

	return &HashBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"哈希函数测试（Hash Functions）",
			"测试各类哈希函数性能",
			"加密性能",
			testFunc,
			100000, // 10万次哈希操作
		),
	}
}

func hashTest(operations int) {
	data := make([]byte, 1024)
	_, _ = rand.Read(data)
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
