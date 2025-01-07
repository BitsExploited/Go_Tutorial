package main

import "fmt"

func main() {
  num := []int{1,2,3,4,5,6,7,8,9,10}

  for i, v := range num {
    fmt.Printf("2 x %d = %d\n", i+1, 2*v)
  }
}
