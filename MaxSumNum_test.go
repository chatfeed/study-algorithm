package main

import (
	"math/rand"
	"testing"
)

var list = generateTestData()

func generateTestData() []int {
	var list []int
	r := rand.New(rand.NewSource(1111))
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10000; i++ {
		if r.Intn(2) == 1 {
			list = append(list, -r.Intn(100))
		} else {
			list = append(list, r.Intn(100))
		}
	}
	return list
}

func Benchmark_maxSumNum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxSumNum1(list)
	}
}
func Benchmark_maxSumNum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxSumNum2(list)
	}
}

func Benchmark_maxSumNum3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxSumNum3(list)
	}
}
func Benchmark_maxSumNum4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxSumNum4(list)
	}
}
