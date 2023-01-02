package main
import "fmt"

func main(){
  var bilangan int
  fmt.Scan(&bilangan)
  fmt.Println(dec2Bin(bilangan))
}

func dec2Bin(n int) string {
  // rekursif
  if n == 0 {
    return "0"
  }else if n == 1 {
    return "1"
  }else {
    if n%2 == 1 {
      return dec2Bin(n/2) + "1"
    }else {
      return dec2Bin(n/2) + "0"
    }
  }
}
