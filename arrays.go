package main

import "fmt"


func testingArrays(){
  var a [2] string
  a[0] = "Hello"
  a[1] = "World"

  fmt.Println(a[0], a[1])
  fmt.Println(a)

  primes := [6] int {2, 3, 5, 7, 11, 13}
  fmt.Println(primes)

  testingSlices()
}

func testingSlices(){
  primes := [6] int {2, 3, 5, 7, 11, 13}
  var s []int = primes[1:4]
  fmt.Println(s)


 // a := make([]int, 5) 
 // a := make([]int, 0, 5)  3rd arg for capacity

  var testSlice []int

  testSlice = append(testSlice, 1)
}
