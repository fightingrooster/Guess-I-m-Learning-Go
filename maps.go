package main


import (
  "fmt"
  "strings"
)


func WordCount(s string) map[string]int {
  //yayaya can be better but good enough
  wordCountMap := make(map[string]int)
  words := strings.Fields(s)

  for _, value := range words{
    _, ok := wordCountMap[value]
    if !ok{
      wordCountMap[value] = 1
    }else{
      wordCountMap[value] += 1
    }
  }
  return wordCountMap
}



type Vertex struct{
  Lat, Long float64
}


var m = map[string]Vertex{
  "K1":{
    40.22, 90.000,
  },
  "K2":{
    41.22, 91.000,
  },
}


func testingMaps(){
  m = make(map[string]Vertex)
  m["test"] = Vertex{
    40.20, -879.00,
  }

  value, ok := m["test"]

  if ok {
    fmt.Println("Key is present", ok)
    fmt.Println("Value", value)
  }else{
    fmt.Println("Key is no present")
  }
}

