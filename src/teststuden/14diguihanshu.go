package main

import(
  "fmt"
)

func sums( x int ) ( sumR int) {
  if (x > 0) {
    sumR = x + sums( x - 1 )    //连加 回调
  }
  return sumR

}

func main() {
  var a int = 10
  b := sums( a )
  fmt.Println( b )

  var i int
  for i = 0 ; i < 20; i++{
    fmt.Println( fbnqsl( i ) )
  }

}


//斐波那契数列


func fbnqsl( x int ) int{

  if x < 2{
    return x
  }
  return fbnqsl( x - 2 ) + fbnqsl( x - 1 )

}
