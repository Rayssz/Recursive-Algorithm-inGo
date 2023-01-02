package main
import "fmt"

const NMAX = 1000
type arrInt [NMAX]int

func main(){
  var A arrInt
  var i, N, X int

  fmt.Scan(&N, &X)
  for i = 0; i<N; i++ {
    fmt.Scan(&A[i])
  }
  fmt.Println(binarySearch(A,N,X))
}

func binarySearch(T arrInt, n,X int) int{
  //fungsi master
  return search(T,0,n-1,X)
}

func search(T arrInt, left, right, X int) int {
  //fungsi rekursif
  var mid int = (left + right) / 2
  if left > right {
    return -1
  }else if X == T[mid] {
    return mid
  }else {
    if X > T[mid] {
      return search(T, mid+1, right, X)
    }else {
      return search(T, left, mid-1, X)
    }
  }
}
