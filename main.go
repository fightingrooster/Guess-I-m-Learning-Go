package main

import(
  "fmt"
  "math/cmplx"
  "time"
)

func main(){
  fmt.Println("Testing add:",add(5, 10))

  a,b := swap("hello", "world")
  fmt.Println("Testing swap:",a, b) 

  fmt.Println(split(22))

  var i int
  fmt.Println(i,"this is a var")

  var (
	  ToBe   bool       = false
	  MaxInt uint64     = 1<<64 - 1
	  z      complex128 = cmplx.Sqrt(-5 + 12i)
  )

  fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

  sum := 0
  for i:= 0; i< 10; i++{
    sum +=i
  }
  fmt.Println("sum: ", sum)
  
  for sum < 1000{
    sum += sum
  }

  fmt.Println("sum: ", sum)


  fmt.Println(sqrt(2), sqrt(-4))

  t := time.Now()
  // -- this is interesting
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

  testPointers()  
  testingArrays()
  testingMaps()
  fmt.Println("Result: " ,WordCount("I am The Best Because I am Good"))
}


/*
--- Basic types
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/
