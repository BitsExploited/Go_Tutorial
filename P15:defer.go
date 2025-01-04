package main

import "fmt"

func main(){
  defer fmt.Println("Cracker")

  fmt.Println("Hash ")
}

// This the defer function is run immediately but the funciton is called only after the nearby functions are called and executes at the last
