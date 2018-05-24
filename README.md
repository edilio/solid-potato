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
