package main

type WeighedQuickUnion struct {
	arr  []int
	size []int
}

func createWeighedQuickUnion(n int) WeighedQuickUnion {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	size := make([]int, n)
	for i := 0; i < n; i++ {
		size[i] = 1
	}
	return WeighedQuickUnion{arr: arr, size: size}
}

func (q WeighedQuickUnion) connect(a, b int) {
	if q.size[a] < q.size[b] {
		q.arr[b] = a
		q.size[b] += q.size[a]
	} else {
		q.arr[b] = a
		q.size[a] += q.size[b]
	}
}

func (q WeighedQuickUnion) isConnected(a, b int) bool {
	for q.arr[a] != a {
		a = q.arr[a]
	}

	for q.arr[b] != b {
		b = q.arr[b]
	}

	return a == b
}
