package main

// ComputeBenchmark 计算密集型测试（基于Pi计算）
type ComputeBenchmark struct {
	*BaseBenchmark
}

// NewComputeBenchmark 创建计算密集型测试实例
func NewComputeBenchmark() *ComputeBenchmark {
	testFunc := func(workload int) {
		n := workload // n代表Pi计算的位数
		_, _, _ = computePi(n)
	}

	return &ComputeBenchmark{
		BaseBenchmark: NewBaseBenchmark(
			"圆周率计算（Pi Calculation）",
			"计算圆周率来测试CPU整数运算性能",
			"计算密集型",
			testFunc,
			10000, // 计算1万位Pi
		),
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
