package main

type QuickFindUF []int

func createQuickFindUF(n int) QuickFindUF {
	arr := make(QuickFindUF, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func (arr QuickFindUF) connect(a, b int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == a {
			arr[i] = b
		}
	}
}

func (arr QuickFindUF) isConnected(a, b int) bool {
	return arr[a] == arr[b]
}
