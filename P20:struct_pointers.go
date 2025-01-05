package main

import (
	"fmt"
)

type two_num struct {
  A int 
  B int
}

func main(){
  a := two_num{42, 24}
  p := &a
  p.B = 72 // Accessing and changing the value of the struct field of 'B'
          // Struct fields can be accessed through dot.
  fmt.Println(a)
}
