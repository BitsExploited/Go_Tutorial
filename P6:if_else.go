package main
import "fmt"

func main(){
  if 8 % 2 == 0{
    fmt.Println("8 is a odd number")
  }

  n := 5
  if n < 0{
    fmt.Println(n, "is a negative number")
  } else if n < 10 && n > 0{
    fmt.Println(n, "is a positive number and is a single digit number")
  }else {
    fmt.Println(n ,"has multiple digits")
  }
}
