package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	time1 := time.Now()
	data, err := json.Marshal(time1)
	if err != nil {
		log.Fatal(err)
	}

	var time2 time.Time
	if err := json.Unmarshal(data, &time2); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Is time1 == time2 :", time1 == time2) //prints false

	fmt.Println("Is time1.Equal(time2) :", time1.Equal(time2)) //prints false

	fmt.Printf("n1 is %T, n2 is %T\n", time1, time2)

}

// Explanation:
// "Operating systems offer two distinct types of clocks: the 'wall clock,' which can be adjusted for
// synchronization purposes, and the 'monotonic clock,' which remains consistent.

// Monotonic clocks serve the purpose of measuring time spans and are designed to mitigate issues that might arise,
// such as the disruption caused by transitions to daylight saving time on a computer.
// The individual value of a monotonic clock lacks significance on its own; rather, its utility lies
// in the calculation of the time elapsed between two separate readings of the monotonic clock.

// During the comparison of time.Time instances using the equality operator == in Go, the comparison
// encompasses the various fields within the time.Time structure, including the monotonic clock reading.
// However, when Go serializes a time.Time instance to JSON format, it omits the monotonic clock reading
// from the resulting output. Consequently, upon deserialization of the JSON data
// into a time.Time instance (time2), the monotonic reading is absent,
// causing the comparison to yield an inaccurate result.

// in time.Time’s documentation:
// “In general, prefer t.Equal(u) to t == u, since t.Equal uses the most accurate comparison available
// and correctly handles the case when only one of its arguments has a monotonic clock reading.”
