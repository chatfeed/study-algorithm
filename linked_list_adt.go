package main

type List struct {
	data int
	next int
}

func main() {
	var bucket []List
	n := 11
	var bucket []List
	for i := 0; i < n; i++ {
		bucket[i].data = 1
		bucket[i].next = 2
	}
}
