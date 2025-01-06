package main

import "fmt"

func main() {
  list := [10]int{1,2,3,4,5,6,7,8,9,10}

  var slice []int = list[1 : 4]
    fmt.Println(slice)
}
