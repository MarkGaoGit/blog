package main

import(
  "fmt"
)

func main (){
  var arr = []int { 1,2,3,4,5,6,7,8,9,10}
  arrAvg := test( arr, len(arr) )
  fmt.Println( arrAvg )


}


//把数组穿入 然后求平均值
func test( arr[]int, size int  ) float64 {

  var i, sum int
  var avg float64

  for i = 0; i < size; i++ {
      sum += arr[i]
  }
  avg = float64( sum / size )

  return avg


}
