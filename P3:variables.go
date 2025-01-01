package main
import "fmt"


func main(){
  var a = "String" // Variable declaration for a String
  fmt.Println(a) // This syntax can also be done outside the function

  b := "Hash-Cracker" // := assings the datatype of the variable in the right
  fmt.Println(b) // Can only be used inside the variable

  var x, y int = 45, 36 
  fmt.Println(x, y)

  var n int 
  fmt.Println(n) // Will return 0 since the default value of the int datatype is zero

  e := true
  fmt.Println(e)
}
