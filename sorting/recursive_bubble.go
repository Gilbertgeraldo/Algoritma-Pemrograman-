package sorting

import (

)

func Bubblesort(arr []int) {
	n := len(arr)

	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - 1 - i; j++ {
			if arr[j] > arr[j + 1] {
			arr[j],arr[j + 1] = arr[j + 1],arr[j]
			}
		}
	}
}

func RecBubbleSOrt(arr []int, n int) {
	if n == 1 { 
		return
	}

	for i := 0; i < n - 1; i++ {
		if arr[i] > arr[i + 1] {
			arr[i],arr[i + 1] = arr[i + 1],arr[i]
		}
	}
	RecBubbleSOrt(arr, n -1)
}

func InsertionSort(arr []int) {
	for i := 1; i < len(arr);i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j + 1] = arr[j]
			j--
		}
		
		arr[j + 1] = key
	}
}

func SelectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIndex := i

		for j := i+1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i],arr[minIndex] = arr[minIndex],arr[i]
	}
}


