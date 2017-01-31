package slices

// https://blog.golang.org/slices

import "fmt"

type tester interface {
    Tester()
}

func main() {
    Tester()
}

func Tester() {
     var buffer [256]byte
    slice := buffer[10:20]
    for i := 0; i < len(slice); i++ {
        slice[i] = byte(i)
    }
    fmt.Println("before", slice)
    AddOneToEachElement(slice)
    fmt.Println("after", slice)
}

func AddOneToEachElement(slice []byte) {
    for i := range slice {
        slice[i]++
    }
}

