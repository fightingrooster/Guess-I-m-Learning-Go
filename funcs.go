package main

import (
  "fmt"
  "math"
)

//-- can also be x, y int
func add(x int, y int) int{
  return  x + y
}

func swap(x, y string)(string, string){
  return y, x
}

//-- naked return
func split(sum int)(x, y int){
  x = sum * 4 /9
  y = sum - x
  return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func testPointers(){
  i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
