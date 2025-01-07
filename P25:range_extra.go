package main

import "fmt"

func main() {
  list := []int{9, 18, 27, 36, 45}

  for _, v := range list { // In case if we do not want the scope of the program we can skip it using "_"
    fmt.Println(v)
  }
}
