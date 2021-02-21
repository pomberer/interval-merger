package main

import (
	"math/rand"
	"testing"
)

func TestMerge(t *testing.T) {
	
	interval1 := Interval{Low: 25, High: 30}
	interval2 := Interval{Low: 2, High: 19}
	interval3 := Interval{Low: 14, High: 23}
	interval4 := Interval{Low: 4, High: 8}

	intervals := []Interval{interval1,interval2,interval3,interval4}
	result := merge(intervals)
	if len(result) != 2 {
		t.Errorf("test merge failed, expected %v results got %v results",2, len(result))
		t.FailNow()
	}

	if result[0].Low != 2 || result[0].High != 23 {
		t.Errorf("test merge failed, expected low interval %v,%v got %v,%v",2, 23,result[0].Low,result[0].High)
		t.FailNow()
	}

	if result[1].Low != 25 || result[1].High != 30 {
		t.Errorf("test merge failed, expected low interval %v,%v got %v,%v",25, 30,result[1].Low,result[1].High)
		t.FailNow()
	}

}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
	var intervals  []Interval
	for i :=0; i< 10000;i++{
		low := rand.Intn(100)
		high := rand.Intn(100-low+1)+low
		intervals = append(intervals, Interval{Low: low, High: high} )
	}
		merge(intervals)
	}
}
