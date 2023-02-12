package main

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
