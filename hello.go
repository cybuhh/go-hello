package main

import (
    "fmt"
)

const (
    message2 = "Hello, Go!"
)

var (
    message = message2

    myArray = [...]int {11, 22, 33}
    myArray2 = [3]int{}
    mySlice = myArray[:]
    mySlice2 = []float32{12., 15., 18.}

    mySlice3 = make([]float32, 10)

    myMap = make(map[int]string)
)

func main() {
    println(message);
    fmt.Println(message2)
    fmt.Println(myArray)

    myArray2[0] = 00
    myArray2[1] = 11
    myArray2[2] = 22
    fmt.Println(myArray2)

    fmt.Println(mySlice)
    mySlice = append(mySlice, 100)
    fmt.Println(mySlice)
    fmt.Println(mySlice2)

    mySlice3[0] = 12.
    mySlice3[1] = 15.
    mySlice3[2] = 18.
    fmt.Println(mySlice3)

    myMap[42] = "Foo"
    myMap[12] = "Bar"
    fmt.Println(myMap)
    fmt.Println(myMap[99])
    fmt.Println(myMap[42])

    fmt.Println("slice", myArray2[0:2])
    fmt.Println("slice", mySlice3[:2])
    loops()
    branches()
    readStdid()

    param := "ref foo"
    fmt.Println(param)
    reference(&param)
    fmt.Println(param)

    variadicFunction("hello", "my", "ugly", "friend")

    numTerms, sum := returnValuesFunc(1, 3, 5 ,9)
    println(numTerms, sum)

    numTerms2, sum2 := returnNamedValuesFunc(1, 3, 5 ,9)
    println(numTerms2, sum2)
}

func readStdid() {
    var option string
    fmt.Println("Please choose an option: ")
    fmt.Scanln(&option)
}

func init() {
    msg2 := "fo"
    println(msg2)
    message = "Hello, Ugly"
}


func branches() {
    foo := 1

    if foo == 1 {
        println("bar")
    }

    if foo == 0 {
        println("bar")
    } else {
        println("baz")
    }

    switch {
        case foo == 0: {
            println("bar-0")
        }
        case foo == 1: {
            println("bar-1")
        }
    }

    switch foo {
        case 0: {
            println("bar-0")
        }
        case 1: {
            println("bar-1")
        }
    }


}

func loops() {
    for i:=0; i < 5; i++ {
        println(i)
    }

    s := []string{"foo", "bar", "buzdaje ok"}

    for idx, v := range s {
        println(idx, v)
    }

    m := make(map[string]string)
    m["first"] = "foo"
    m["second"] = "bar"
    m["thrid"] = "buz"

    for k, v := range m {
        println(k, v)
    }
}

func reference(param *string) {
    *param = "ref bar"
}

func variadicFunction(params ...string) {
    for _, param := range params {
        println(param)
    }
}

func returnValuesFunc(terms ...int) (int, int) {
    result := 0
    for _, term := range terms {
        result += term
    }

    return len(terms), result
}

func returnNamedValuesFunc(terms ...int) (numTerms int, sum int) {
    for _, term := range terms {
        sum += term
    }

    numTerms = len(terms)

    return
}
