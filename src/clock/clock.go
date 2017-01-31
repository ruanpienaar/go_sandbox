package main
import (
    "fmt"
    "strconv"
    "os"
)
const testVersion = 4
type Clock struct{ h, m int }

func main() {
    fmt.Println(os.Args[1])
    fmt.Println(os.Args[2])
    h64, _ := strconv.ParseInt(os.Args[1], 10, 64)
    m64, _ := strconv.ParseInt(os.Args[2], 10, 64)
    h := int(h64)
    m := int(m64)
    C := New(h, m)
    fmt.Println(C.String())
    // fmt.Println("10:03")
    // fmt.Println(New(10, 0).Add(3))
}

func New(hour, minute int) Clock {
    var totalmins int = totalmins(hour, minute)
    var h int = 0
    h2, totalmins2 := check_floor(h, totalmins)
    h3, totalmins3 := check_ceiling(h2, totalmins2)
    return Clock{h:h3, m:totalmins3}
}

func check_floor(h, mins int) (int, int) {
    for mins < 0 {

        if h <= 0 {
            h = 23
        } else {
            h = h - 1
        }

        mins += 60
    }
    return h, mins
}
func check_ceiling(h, mins int) (int, int) {
    for mins >= 60 {
        h = h + 1
        if h >= 24 {
            h = 0
        }
        mins -= 60
    }
    return h, mins
}

func totalmins(h, m int) int {
    var hours_in_mins int = h*60
    var totalmins int = hours_in_mins + m
    return totalmins
}

func (c Clock) String() string {
    return fmt.Sprintf("%s:%s", c.HStr(), c.MStr())
}

func (c Clock) Add(minutes int) Clock {
    h2, totalmins2 := check_floor(c.H(), c.M()+minutes)
    h3, totalmins3 := check_ceiling(h2, totalmins2)
    return Clock{h:h3, m:totalmins3}
}

func (c Clock) H() int {
    return c.h
}

func (c Clock) M() int {
    return c.m
}

func (c Clock) HStr() string {
    if c.h > 9 {
        return strconv.Itoa(c.h)
    } else {
        return "0"+strconv.Itoa(c.h)
    }
}

func (c Clock) MStr() string {
    if c.m > 9 {
        return strconv.Itoa(c.m)
    } else {
        return "0"+strconv.Itoa(c.m)
    }
}