package main

import "fmt"

func main(){
  var a, b int = 6, 7
  var c float64 = float64(a) // it will be converted to float64 only if its converted explicitly
  var d uint = uint(a)
  fmt.Println(a, b, c, d)
}
