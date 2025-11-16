package main

import (
	"math"
	"math/rand"
)

// FloatBenchmark 浮点运算性能测试
type FloatBenchmark struct {
	*BaseBenchmark
}

// NewFloatBenchmark 创建浮点运算测试实例
func NewFloatBenchmark() *FloatBenchmark {
	testFunc := func(workload int) {
		floatTest(workload)
	}

	return &FloatBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"浮点运算测试（Floating Point）",
			"测试浮点数运算性能",
			"浮点性能",
			testFunc,
			5000000, // 500万次浮点运算
		),
	}
}

func floatTest(operations int) {
	result := 0.0
	r := rand.New(rand.NewSource(42))
	for i := 0; i < operations/8; i++ {
		a := r.Float64()
		b := r.Float64()
		result += a + b
		result += a - b
		result += a * b
		if b != 0.1 && b != 0 {
			result += a / b
		}
	}
	for i := 0; i < operations/8; i++ {
		x := r.Float64()*10 + 1
		y := r.Float64()*5 + 0.5
		result += math.Pow(x, y)
		result += math.Sqrt(x)
		result += math.Cbrt(x)
		result += math.Log10(x)
	}
	for i := 0; i < operations/5; i++ {
		x := r.Float64() * 100
		result += math.Sin(x)
		result += math.Cos(x)
		result += math.Sqrt(math.Abs(x))
		result += math.Log(math.Abs(x) + 1)
		result += math.Exp(-math.Abs(x) / 10)
	}
	for i := 0; i < operations/5; i++ {
		x := r.Float64() * 10
		y := r.Float64() * 10
		result += math.Pow(x, y)
		result += math.Atan2(x, y)
		result += math.Hypot(x, y)
	}
	for i := 0; i < operations/10; i++ {
		a11, a12, a21, a22 := r.Float64(), r.Float64(), r.Float64(), r.Float64()
		b11, b12, b21, b22 := r.Float64(), r.Float64(), r.Float64(), r.Float64()
		c11 := a11*b11 + a12*b21
		c12 := a11*b12 + a12*b22
		c21 := a21*b11 + a22*b21
		c22 := a21*b12 + a22*b22
		result += c11 + c12 + c21 + c22
	}
	for i := 0; i < operations/10; i++ {
		angle := r.Float64() * 2 * math.Pi
		result += math.Tan(angle)
		result += math.Asin(r.Float64())
		result += math.Acos(r.Float64())
	}
	_ = result
}

// MatrixBenchmark 矩阵运算测试
type MatrixBenchmark struct {
	*BaseBenchmark
}

func NewMatrixBenchmark() *MatrixBenchmark {
	testFunc := func(workload int) {
		// 将workload映射为矩阵大小，使工作量合理
		size := 200
		if workload > 200 {
			// 工作量增加时，增大矩阵大小（立方根关系，因为矩阵乘法是O(n³)）
			ratio := float64(workload) / 200.0
			size = int(200.0 * pow(ratio, 1.0/3.0))
		}
		matrixTest(size)
	}

	return &MatrixBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"矩阵运算测试（Matrix Operations）",
			"测试矩阵运算性能",
			"浮点性能",
			testFunc,
			200, // 基础矩阵大小200x200
		),
	}
}

// 简单的pow函数实现，避免引入math包
func pow(x, n float64) float64 {
	if n == 0 {
		return 1
	}
	result := 1.0
	for i := 0; i < int(n); i++ {
		result *= x
	}
	return result
}

func matrixTest(size int) {
	a := make([][]float64, size)
	b := make([][]float64, size)
	c := make([][]float64, size)
	for i := 0; i < size; i++ {
		a[i] = make([]float64, size)
		b[i] = make([]float64, size)
		c[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			a[i][j] = float64(i*j) + 0.1
			b[i][j] = float64(i+j) + 0.2
		}
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
}
