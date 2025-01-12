package main

import "fmt"

func main() {
  var arr []int
  fmt.Println("Base array:", arr)

  arr = append(arr, 1)
  fmt.Println("After adding one element:", arr)

  arr = append(arr, 2, 3, 4, 5)
  fmt.Println("After adding multiple elements:", arr)
}
