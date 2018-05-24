package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func insertionSort(a []int) {
	n := len(a)
	for i := 1; i < n; i++ {
		for k := i; k > 0 && a[k] < a[k-1]; k-- {
			a[k-1], a[k] = a[k], a[k-1]
		}
	}
}

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

func bubbleSort1(a []int) {
	n, swapped := len(a), true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if a[i-1] > a[i] {
				a[i], a[i-1] = a[i-1], a[i]
				swapped = true
				// invariant: a[1..i] in final position
			}
		}
	}
}

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

type sortFunc = func(a []int)

func fromStringToIntArr(s string) []int {
	ret := []int{}
	lst := strings.Split(s, ",")
	for _, str := range lst {
		x, err := strconv.Atoi(str)
		if err == nil {
			ret = append(ret, x)
		}
	}
	return ret
}

func main() {
	algos := map[string]sortFunc{
		"Quicksort": quicksort,
		"Insertion": insertionSort,
		"Bubble":    bubbleSort,
	}

	http.HandleFunc("/sort", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			respMsg := msg
			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
			s := strings.Split(string(msg), ":")
			if len(s) == 2 {
				sortAlgo, ok := algos[s[0]]
				if ok {
					arr := fromStringToIntArr(s[1])
					sortAlgo(arr)
					sorted := fmt.Sprintf("%v", arr)
					respMsg = []byte(sorted)
				}
			}

			// Write message back to browser
			if err = conn.WriteMessage(msgType, respMsg); err != nil {
				return
			}
		}

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}
