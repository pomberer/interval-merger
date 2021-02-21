package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting Interval Merger")
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	log.Println("Interval Merger Ready")
	api.HandleFunc("/merge", post).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func post(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var intervals IntervalList
	_ = json.NewDecoder(request.Body).Decode(&intervals)
	json.NewEncoder(writer).Encode(IntervalList{merge(intervals.Intervals)})
}

func merge(intervals []Interval) []Interval {
	start := time.Now()

	if len(intervals) <= 2 {
		return intervals
	}

	sort.Sort(IntervalList{intervals})
	merged := make([]Interval, 0)
	a := &intervals[0]
	listLength := len(intervals)

	for i := 1; i < listLength; i++ {
		b := &intervals[i]
		if a.High >= b.Low {
			if a.High < b.High {
				a.High = b.High
			}
		} else {
			merged = append(merged, *a)
			a = b
		}
	}

	merged = append(merged, *a)
	elapsed := time.Since(start)
	log.Printf("merge for %v intervals took %s", len(intervals) , elapsed)
	return merged

}

type Interval struct {
	Low, High int
}

type IntervalList struct {
	Intervals []Interval `json:"intervals"`
}

//Len, Less, Swap needed to execute sort in go
func (list IntervalList) Len() int {
	return len(list.Intervals)
}
func (list IntervalList) Less(a, b int) bool {
	return list.Intervals[a].Low < list.Intervals[b].Low
}

func (list IntervalList) Swap(a, b int) {
	list.Intervals[a], list.Intervals[b] = list.Intervals[b], list.Intervals[a]
}
