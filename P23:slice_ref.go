package main
import "fmt"

func main() {
  list := [5]int{9, 18, 27, 36, 45}
  fmt.Println(list)

  a := list[0 : 3]
  b := list[3 : 5]
  fmt.Println(a, b)

  b[0] = 90
  fmt.Println(a, b)
  fmt.Println(list)
}
