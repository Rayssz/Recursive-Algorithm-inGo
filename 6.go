package main
import "fmt"

const NMAX = 1000
type arrInt [NMAX]int

func main (){
  var A arrInt
  var i, N int

  fmt.Scan(&N)
  for i=0; i<N; i++ {
    fmt.Scan(&A[i])
  }
  fmt.Println(check(A,0,N-1,N))
}

func check(T arrInt, i,j,n int) bool {
  //fungsi rekursif
  //base case
  if i == n/2 {
    return T[i] == T[j]
  }else {
    return (T[i] == T[j]) && check(T, i+1, j-1, n)
  }
}
