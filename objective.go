package main

import "fmt"

func main() {
    foo0 := myStruct{"bar"}
    fmt.Println(foo0)

    foo1 := new(myStruct)
    foo1.myField = "bar"
    fmt.Println(foo1)

    // create class by reference
    foo2 := &myStruct{"bar"}
    fmt.Println(foo2)

    foo3 := newMyStruct()
    foo3.myMap["bar"] ="baz"
    fmt.Println(foo3)

    mp := messagePrinter{"messagePrinter foo"}
    mp.printMessage()

    mp2 := enhancedMessagePrinter{}
    mp2.message = "enhancedMessagePrinter foo"
    mp2.printMessage()
}

type myStruct struct {
    myField string
}

type myStruct2 struct {
    myMap map[string]string
}

func newMyStruct() *myStruct2 {
    result := myStruct2{}
    result.myMap = map[string]string{}

    return &result
}

type messagePrinter struct {
    message string
}

func (mp *messagePrinter) printMessage() {
    println(mp.message)
}

type enhancedMessagePrinter struct {
    messagePrinter
}

