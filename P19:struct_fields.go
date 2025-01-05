package main

import "fmt"

type two_num struct {
  A int 
  B int
}

func main(){
  a := two_num{42, 24}
  a.B = 36 // Accessing and changing the value of the struct field of 'B'
  // Struct fields can be accessed through dot.
  fmt.Println(a.B)
}
