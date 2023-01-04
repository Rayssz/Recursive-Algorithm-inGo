package main
import "fmt"

const NMAX = 1000
type arrInt [NMAX]int

func main (){
  var A arrInt
  var i, N int

  fmt.Scan(&N)
  for i=0 ; i<N; i++ {
    fmt.Scan(&A[i])
  }
  fmt.Println(sumArray(A, 0, N))
}

func sumArray(T arrInt, i,n int) int {
  //fungsi rekursif
  //base case
  if i == n {
    return 0
  }else {
    return T[i] + sumArray(T, i+1, n)
  }
}
