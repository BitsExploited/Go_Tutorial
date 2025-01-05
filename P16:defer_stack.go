package main

import "fmt"

func main(){
  fmt.Println("Welcome to the Loop")

  for i := 0; i <= 5; i++ {
    defer fmt.Println(i)
  }

  fmt.Println("This statement will execute first")
}
