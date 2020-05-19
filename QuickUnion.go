package main

type QuickUnion []int

func createQuickUnion(n int) QuickUnion {
	arr := make(QuickUnion, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func (q QuickUnion) connect(a, b int) {
	q[b] = a
}

func (q QuickUnion) isConnected(a, b int) bool {
	// 0 1 2 3 4 5 6 7 8 9
	// 1 8 1 8 3 0 5 1 8 8
	for q[a] != a {
		a = q[a]
	}

	for q[b] != b {
		b = q[b]
	}

	return a == b
}
