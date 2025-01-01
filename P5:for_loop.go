package main
import "fmt"

func main(){

  fmt.Println("First way of declaring for loop")
  for j := 0; j <= 15;{
    fmt.Println(j)
    j += 1
  }

  fmt.Println("Second way of declaring for loop")
  for j := 0; j <= 15; j++{
    fmt.Println(j)
  }

  fmt.Println("Third way of declaring for loop")
  for i := range 15{
    fmt.Println(i)
  }

  fmt.Println("Fourth way of declaring for loop")
  for i := range 15{
    if i % 2 == 0{
      continue
    }
    fmt.Println(i)
  }
}
