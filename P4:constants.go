package main
import (
  "fmt"
  "math"
)

const s string = "String" // Similar to that of "var", "const" can also be decalred outside a function

func main(){
  fmt.Println(s)

  const n = 6000000

  const d = 3e20 / n
  fmt.Println(d)

  fmt.Println(int64(d)) // Stores the value of "d" in 64-bit

  fmt.Println(math.Sin(n))

}
