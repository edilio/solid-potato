# solid-potato
Sorting implementations tested with websockets(just for fun)

## Sorting algos
* insertion
* bubble 
* quicksort

## To test it
go get github.com/gorilla/websocket

then

go run server.go

browse to http://localhost:8080

Use comma separated integers without spaces in order to test 

### insertion sort code
```golang
func insertionSort(a []int) {
	n := len(a)
	for i := 1; i < n; i++ {
		for k := i; k > 0 && a[k] < a[k-1]; k-- {
			a[k-1], a[k] = a[k], a[k-1]
		}
	}
}
```

### bubble sort(fast implementation)
```golang
func bubbleSort(a []int) {
	n := len(a)
	for n != 0 {
		newn := 0
		for i := 1; i < n; i++ {
			if a[i-1] > a[i] {
				a[i], a[i-1] = a[i-1], a[i]
				newn = i
				// invariant: a[1..i] in final position
			}
			fmt.Println(i, newn)
		}
		n = newn
	}
}
```

### quicksort
```golang
// select median between a[lo], a[hi] and a[middle]
func pivot(a []int, lo, hi int) int {
	first, last := a[lo], a[hi]
	middleIndex := lo + (hi-lo)/2

	middle := a[middleIndex]
	// fmt.Printf("Pivot Index: %d Median: %d\n", middleIndex, middle)
	// fmt.Printf("first: %v middle: %v last: %v\n", first, middle, last)

	if (first <= middle && middle <= last) || (last <= middle && middle <= first) {
		return middle
	} else if (middle <= first && first <= last) || (last <= first && first <= middle) {
		return first
	} else if (middle <= last && last <= first) || (first <= last && last <= middle) {
		return last
	}

	return last
}

func partition(a []int, p, lo, hi int) (int, int) {
	leftmark := lo + 1
	rightmark := hi

	done := false
	for !done {
		for leftmark <= rightmark && a[leftmark] <= p {
			leftmark++
		}

		for a[rightmark] >= p && rightmark >= leftmark {
			rightmark--
		}

		if rightmark < leftmark {
			done = true
		} else {
			a[leftmark], a[rightmark] = a[rightmark], a[leftmark]
		}
	}
	a[lo], a[rightmark] = a[rightmark], a[lo]

	return leftmark, rightmark
}

func _quicksort(a []int, lo, hi int) {

	if lo < hi {
		p := pivot(a, lo, hi)
		left, right := partition(a, p, lo, hi) // note: multiple return values
		_quicksort(a, lo, left-1)
		_quicksort(a, right+1, hi)
	}
}

func quicksort(a []int) {
	_quicksort(a, 0, len(a)-1)
}
```
