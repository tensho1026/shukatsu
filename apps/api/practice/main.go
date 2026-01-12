package main

import "fmt"


func add(x,y int)int {
return x + y
}

func main () {
	type Vertex struct {
		x int
		y int
		name string

	}

	var a Vertex

	a.x = 1
	a.y = 2
	a.name = "たかし"

	fmt.Print(a)
// var a [3]int
// a[0] = 1
// a[1] = 1
// a[2] = 1

// println(a)

	// println("Hello World")

	// var a string = "1"
	// println(a)
	// b := "a"
	// println(b)

	// const c = "aa"

	// println(add(1,2))

// 	for i :=0;i<3;i++ {
// 		println(i)
// 	}

// 	var aaaa int = 10

// 	if aaaa < 1 {
// 	println("true")
//  } else if aaaa > 1{
// 	println("false")
//  }
}
