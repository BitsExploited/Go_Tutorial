package main
import "fmt"

func main(){
  b := [5]int{1,2,3,4,5}
  fmt.Println(b)

  b = [...]int{10,20,30,40,50} // when specified [...] the compiler counts the number elements instead of the user specifying the number of elements
  fmt.Println(b)
  
  b = [...]int{15, 3: 20, 67} // Specifying 3 : 20 will place the number 20 in the index 3
  fmt.Println(b)

  var arr = [2][3]int{ // It is necessary to initialise the variable before storing an array unlike other datatypes
    {1,2,3},
    {4,5,6},
  }
  fmt.Println(arr)

}
