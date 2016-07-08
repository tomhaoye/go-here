package main

import "fmt"
import "container/list"

func main() {
        bar(6)
}

func bar(i int) {
        l := list.New()
        for i := 0; i < 9; i++ {
                l.PushBack(i)
        }
        l.PushBack(6)

        for e := l.Front(); e != nil; e = e.Next() {
                fmt.Print(e.Value.(int))
        }

        var next *list.Element
        for e := l.Front(); e != nil; {
                if e.Value.(int) == i {
                        next = e.Next()
                        l.Remove(e)
                        e = next
                } else {
                        e = e.Next()
                }
        }

        fmt.Printf("\n")
        for e := l.Front(); e != nil; e = e.Next() {
                fmt.Print(e.Value.(int))
        }
        fmt.Printf("\n")
}