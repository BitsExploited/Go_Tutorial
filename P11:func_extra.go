package main
import "fmt"

func add(x, y int) int{ // instead of intx, int y it can be shrinked to x, y int
	return x + y
}

func main(){
	fmt.Println(add(45, 36))
}
