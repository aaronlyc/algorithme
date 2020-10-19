package sort7

func heapSort(data []int, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		swap(data, first, first+i)
		//fmt.Printf("sorted swap the num %d and %d\n", data[first], data[first+i])
		siftDown(data, lo, i, first)
	}
}

func siftDown(data []int, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(data, first+child, first+child+1) {
			child++
		}
		if !less(data, first+root, first+child) {
			return
		}
		swap(data, first+root, first+child)
		//fmt.Printf("build heap swap the num %d and %d\n", data[first+root], data[first+child])
		root = child
	}
}
