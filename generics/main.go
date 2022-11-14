package main

import "fmt"

type number interface {
    int64 | float64
}

func main() {
    ints := map[string]int64{
        "first": 34,
        "second": 12,
    }

    floats := map[string]float64{
        "first": 35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SunFloats(floats))

    fmt.Printf("Generic sums: %v and %v\n",
        SumIntsOrFloats(ints),
        SumIntsOrFloats(floats))
}

func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

func SunFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}

// Generic Function, transform last two functions in one
// change `v int64 | float64` by interface Number - clean code
func SumIntsOrFloats[k comparable, v number](m map[k]v) v {
    var s v
    for _, v := range m {
        s += v
    }
    return s
}
