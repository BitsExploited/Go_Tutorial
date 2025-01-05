package main

import "fmt"

func main(){
  a := 78

  p := &a
  fmt.Println("The memory address of 'a' is:", p) // Prints the memory address of 'a' 
  fmt.Println("The value of 'a' through the pointer 'p':",*p) // Printing the value of 'a' through p

  *p = 50 // Changing the value of 'a' through the pointer p
  fmt.Println("The value of 'a' after changing the value of 'a' through pointer 'p':",a) // Now printing the value of 'a'
   
}
