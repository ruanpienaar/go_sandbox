package main
import "fmt"
func main() {
    var T int 
    fmt.Scan(&T)
    //fmt.Println("Going to loop now...")
    for i := 1; i < 11; i++ {
        fmt.Printf("%d x %d = %d\n", T, i, T * i)
    }
}