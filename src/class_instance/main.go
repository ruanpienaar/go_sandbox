package main
import "fmt"

type person struct {
    age int
}

func (p person) NewPerson(initialAge int) person {
    if initialAge > 0 {
        p.age = initialAge
    } else {
        fmt.Println("Age is not valid, setting age to 0.")
        p.age = 0
    }
    return p
}

func (p person) amIOld() {
    if p.age < 13 {
        fmt.Println("you are young.");
    } else if p.age >= 13 && p.age < 18 {
        fmt.Println("You are a teenager.");
    } else {
        fmt.Println("You are old.");
    }
}

func (p person) yearPasses() person {
    p.age += 1
    return p
}

func main() {
    var T,age int

    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        if i != 0 {
            fmt.Println()
        }
        fmt.Scan(&age)
        p := person{age: age}
        p = p.NewPerson(age)
        p.amIOld()
        for j := 0; j < 3; j++ {
            p = p.yearPasses()
        }
        p.amIOld()
        //fmt.Println()
    }
}