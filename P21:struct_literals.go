package main

import "fmt"

type rand_num struct {
  X, Y int
}

var (
  a = rand_num{42, 24} // Assigning X as 42 and Y as 24
  b = rand_num{X : 36} // Assigning X as 36 and Y takes the default value of int as 0
  c = rand_num{} // Both X and Y takes the default number
  p = &rand_num{42, 24} // Assinging a pointer value to the struct rand_num
)

func main() {
  fmt.Println(a, b, c, p)
}
