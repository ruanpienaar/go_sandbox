package main
import "fmt"

func main() {
    var l int
    fmt.Scan(&l)
    for i := 0; i < l; i++ {
        var line string
        fmt.Scan(&line)
        var odd, even []string
        for i, letter := range line {
            //fmt.Println(letter)
            if i % 2 == 0 {
                even = append(even, string(letter))
            } else {
                odd = append(odd, string(letter))
            }
        }
        fmt.Println(even, odd)
    }
}