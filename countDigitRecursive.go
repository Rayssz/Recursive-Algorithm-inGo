package main
import "fmt"

func main(){
  var bilangan int
  fmt.Scan(&bilangan)
  fmt.Println(count_digit(bilangan))
}

func count_digit(n int) int {
  // fungsi rekursif
  if n==0{
    return 0
  }else {
    return 1 + count_digit(n/10)
  }
}
