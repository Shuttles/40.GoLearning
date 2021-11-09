package main 

import (
    "fmt"
)



func main() {
    o := 0666
    x := int64(0xdeadbeef)
    ascii := 'a'
    unicode := '国'
    newline := '\n'

    fmt.Printf("%d %[1]o %#[1]o\n", o)
    fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
    fmt.Printf("%d %[1]c %[1]q\n", ascii)
    fmt.Printf("%d %[1]c %[1]q\n", unicode)
    fmt.Printf("%d %[1]c %[1]q\n", newline)

    medals := []string{"gold", "silver", "bronze"}
    for i:= len(medals) - 1; i >= 0; i-- {
        fmt.Println(medals[i])
    }
}
