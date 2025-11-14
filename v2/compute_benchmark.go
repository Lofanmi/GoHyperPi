package main

import (
	"sync"
	"time"
)

// ComputeBenchmark 计算密集型测试（基于Pi计算）
type ComputeBenchmark struct{}

// NewComputeBenchmark 创建计算密集型测试实例
func NewComputeBenchmark() *ComputeBenchmark {
	return &ComputeBenchmark{}
}

func (cb *ComputeBenchmark) Name() string {
	return "圆周率计算（Pi Calculation）"
}

func (cb *ComputeBenchmark) Description() string {
	return "计算圆周率来测试CPU整数运算性能"
}

func (cb *ComputeBenchmark) Category() string {
	return "计算密集型"
}

func (cb *ComputeBenchmark) Run(proc, times int) BenchmarkResult {
	n := 50000 // 默认计算5万位Pi
	p := proc * times
	ch := make(chan float64, p)
	wg := new(sync.WaitGroup)
	wg.Add(p)
	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			_, _, _ = computePi(n)
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
	t1 := float64(n) / single               // 单核性能指标
	tn := float64(p*n) / duration.Seconds() // 多核性能指标
	efficiency := tn / t1 / float64(proc)   // 多核效率
	return BenchmarkResult{
		Name:       cb.Name(),
		Category:   cb.Category(),
		Score:      tn / 1000.0, // 归一化得分
		Duration:   duration,
		SingleRate: t1,
		MultiRate:  tn,
		Efficiency: efficiency,
		Proc:       proc,
		Times:      times,
	}
}

// computePi 计算Pi值（从原代码移植）
func computePi(n int) (i, N int, pi []int) {
	N = n/4 + 3
	pi = make([]int, N)
	var j, k, p, q, r, t, u, v int
	a, b := [2]int{956, 80}, [2]int{57121, 25}
	s := 2
	M := 10000
	e := make([]int, N)
	for {
		s--
		if s+1 == 0 {
			break
		}
		k = s
		e[0] = a[s]
		i = N
		for {
			i--
			if i == 0 {
				break
			}
			e[i] = 0
		}
		q = 1
		for {
			j = i - 1
			if i >= N {
				break
			}
			r, v = 0, 0
			for {
				j += 1
				if j >= N {
					break
				}
				p = r*M + e[j]
				e[j] = p / b[s]
				t = v*M + e[j]
				u = t / q
				r = p % b[s]
				v = t % q
				if k != 0 {
					pi[j] += u
				} else {
					pi[j] -= u
				}
			}
			if e[i] == 0 {
				i++
			}
			q += 2
			if k != 0 {
				k = 0
			} else {
				k = 1
			}
		}
	}
	for {
		i--
		if i == 0 {
			break
		}
		t = pi[i] + s
		pi[i] = t % M
		if pi[i] < 0 {
			pi[i] += M
			s = t/M - 1
		} else {
			s = t / M
		}
	}
	return
}
