package main

import (
	"testing"
)

func main() {

}

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// go test 路径 --coverprofile=输出文件
func TestFibonacci(t *testing.T) {
	fsMap := map[int]int{}
	fsMap[1] = 1
	fsMap[2] = 1
	fsMap[3] = 2
	fsMap[4] = 3
	fsMap[5] = 5
	fsMap[6] = 8
	fsMap[7] = 13
	fsMap[8] = 21
	fsMap[9] = 34

	for k, v := range fsMap {
		fib := Fibonacci(k)
		if v == fib {
			t.Logf("结果正确：n为%d, 值为%d\n", k, fib)
		} else {
			t.Errorf("结果是：期望%d, 但是计算结果%d\n", k, fib)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	b.ResetTimer()   // 重置计时器
	b.ReportAllocs() // 开启内存统计
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}

func BenchmarkFibonacciRunParallel(b *testing.B) {
	n := 10
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Fibonacci(n)
		}
	})
}
