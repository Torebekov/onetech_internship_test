package quicksort

import "math/rand"

func QuickSort(a []int) {
	if len(a) < 2 {
		return
	}
	low, high := 0, len(a)-1
	pivot := rand.Int() % len(a)
	a[pivot], a[high] = a[high], a[pivot]
	for i, _ := range a {
		if a[i] < a[high] {
			a[low], a[i] = a[i], a[low]
			low++
		}
	}
	a[low], a[high] = a[high], a[low]

	QuickSort(a[:low])
	QuickSort(a[low+1:])
}
