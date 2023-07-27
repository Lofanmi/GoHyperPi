package main

import (
	"flag"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/klauspost/cpuid/v2"
)

func main() {
	P := runtime.GOMAXPROCS(0)
	var (
		output         bool
		n, proc, times int
	)
	flag.BoolVar(&output, "output", false, "Output Pi (default false)")
	flag.IntVar(&n, "n", 100000, "Number of Pi")
	flag.IntVar(&proc, "proc", P, "Proc count")
	flag.IntVar(&times, "times", 2, "ComputePi times")
	flag.Parse()

	PrintCPU()

	if output {
		i, N, pi := ComputePi(n)
		PrintPi(i, N, pi)
		return
	}

	p := proc * times
	ch := make(chan float64, p)

	wg := new(sync.WaitGroup)
	wg.Add(p)
	start := time.Now()
	for i := 0; i < p; i++ {
		go func() {
			defer wg.Done()
			t := time.Now()
			_, _, _ = ComputePi(n)
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
	single = single / float64(p*times)

	t1, tn := float64(n)/single, float64(p*n)/duration.Seconds()
	rate := tn / t1

	fmt.Printf("Result:\n[duration:%s] [single-core:%.2f] [multi-core:%.2f] [rate:%.2f]\n", duration, t1, tn, rate)
}

func PrintCPU() {
	fmt.Println("CPU Name:", cpuid.CPU.BrandName)
	fmt.Println("CPU Physical Cores:", cpuid.CPU.PhysicalCores)
	fmt.Println("CPU Threads Per Core:", cpuid.CPU.ThreadsPerCore)
	fmt.Println("CPU Logical Cores:", cpuid.CPU.LogicalCores)
	fmt.Println("CPU Family:", cpuid.CPU.Family, "Model:", cpuid.CPU.Model, "Vendor ID:", cpuid.CPU.VendorID)
	fmt.Println("CPU Features:", strings.Join(cpuid.CPU.FeatureSet(), ","))
	fmt.Println("CPU CacheLine bytes:", cpuid.CPU.CacheLine)
	fmt.Println("CPU L1 Data Cache:", cpuid.CPU.Cache.L1D, "bytes")
	fmt.Println("CPU L1 Instruction Cache:", cpuid.CPU.Cache.L1I, "bytes")
	fmt.Println("CPU L2 Cache:", cpuid.CPU.Cache.L2, "bytes")
	fmt.Println("CPU L3 Cache:", cpuid.CPU.Cache.L3, "bytes")
	fmt.Println("CPU Frequency:", cpuid.CPU.Hz, "Hz")
	fmt.Println("OS:", runtime.GOOS, runtime.GOARCH)
	fmt.Println()
}

func PrintPi(i, N int, pi []int) {
	fmt.Print("3.")
	for i++; i < N-2; i++ {
		s := strconv.Itoa(pi[i])
		s = strings.Repeat("0", 4-len(s)) + s
		fmt.Print(s)
	}
	fmt.Println()
}

func ComputePi(n int) (i, N int, pi []int) {
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
