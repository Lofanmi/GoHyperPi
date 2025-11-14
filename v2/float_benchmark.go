package main

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

// FloatBenchmark 浮点运算性能测试
type FloatBenchmark struct{}

// NewFloatBenchmark 创建浮点运算测试实例
func NewFloatBenchmark() *FloatBenchmark {
	return &FloatBenchmark{}
}

func (fb *FloatBenchmark) Name() string {
	return "浮点运算测试（Floating Point）"
}

func (fb *FloatBenchmark) Description() string {
	return "测试浮点数运算性能"
}

func (fb *FloatBenchmark) Category() string {
	return "浮点性能"
}

func (fb *FloatBenchmark) Run(proc, times int) BenchmarkResult {
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)

	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			floatTest(5000000) // 500万次浮点运算
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

	t1 := 5000000.0 / single                    // 单核浮点运算速率（flops/s）
	tn := float64(p*5000000) / duration.Seconds() // 多核浮点运算速率（flops/s）
	// 避免除零错误
	var efficiency float64
	if t1 > 0 {
		efficiency = tn / t1 / float64(proc)        // 多核效率
	} else {
		efficiency = 0.0
	}

	return BenchmarkResult{
		Name:       fb.Name(),
		Category:   fb.Category(),
		Score:      tn / 1000.0, // 归一化得分
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
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
type MatrixBenchmark struct{}

func NewMatrixBenchmark() *MatrixBenchmark {
	return &MatrixBenchmark{}
}

func (mb *MatrixBenchmark) Name() string {
	return "矩阵运算测试（Matrix Operations）"
}

func (mb *MatrixBenchmark) Description() string {
	return "测试矩阵运算性能"
}

func (mb *MatrixBenchmark) Category() string {
	return "浮点性能"
}

func (mb *MatrixBenchmark) Run(proc, times int) BenchmarkResult {
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)

	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			matrixTest(200) // 200x200矩阵运算
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

	t1 := 1.0 / single                             // 单核矩阵运算速率
	tn := float64(p) / duration.Seconds()          // 多核矩阵运算速率
	// 避免除零错误
	var efficiency float64
	if t1 > 0 {
		efficiency = tn / t1 / float64(proc)          // 多核效率
	} else {
		efficiency = 0.0
	}

	return BenchmarkResult{
		Name:       mb.Name(),
		Category:   mb.Category(),
		Score:      tn * 1000, // 归一化得分
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
	}
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