package main

import "fmt"

func main() {
  a := make([]int, 5) // The first arg is specifying the datatype of the array and the secidn arg is for the size of the slice or array
  fmt.Printf("The array printed is: %v\n", a)

  b := make([]int, 0, 5) // The third arg is for mentioning the capacity of the array
  fmt.Printf("The length of the array is: %d\n", len(b))
  fmt.Printf("The capacity of the array is: %d\n", cap(b))
  fmt.Printf("The resulting array is: %v\n", b)
}
