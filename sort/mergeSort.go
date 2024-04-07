package sort

func merge(a, b []int) []int {
	final := []int{}
	var i, j int

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for i < len(a) {
		final = append(final, a[i])
		i++
	}

	for j < len(b) {
		final = append(final, b[j])
		j++
	}
	return final
}

func mergeSort(A []int) []int {
	if len(A) < 2 {
		return A
	}
	q := len(A) / 2
	a := mergeSort(A[:q])
	b := mergeSort(A[q:])
	return merge(a, b)
}

func mergeIdx(A []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q

	L := make([]int, n1)
	R := make([]int, n2)

	for i := 0; i < n1; i++ {
		L[i] = A[p+i]
	}

	for j := 0; j < n2; j++ {
		R[j] = A[q+j+1]
	}

	i, j := 0, 0

	for k := p; k <= r; k++ {
		if i < n1 && (j >= n2 || L[i] <= R[j]) {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
}

func mergeSortIdx(A []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		mergeSortIdx(A, p, q)
		mergeSortIdx(A, q+1, r)
		mergeIdx(A, p, q, r)
	}
}
