package main

import(
  "fmt"
)

func main (){

  //数组名 长度  类型
  // var arr [10] string

  //初始化数组  数组大小不可以超过 []中设置的长度
  var balance = [5] float64 { 1, 2, 3, 4, 5 }
  fmt.Println( balance )

  //如果[]中不设置大小 则数组自己判断长度  [...]

  var arr = [...] string{"张三", "李四", "王五", "赵六"}
  lengthArr := len(arr)

  balance[4] = 60.5

  var name = arr[2]
  fmt.Println( name )

  fmt.Println( balance )
  fmt.Println( arr )
  fmt.Println( lengthArr )


  // var numberArr = [...] float64
  var numberArr[122] int
  var i, j int

  for i = 0; i < 120; i++ {
    numberArr[ i ] = i+1
  }

  for j = 0; j < len(numberArr); j++ {
    //打印数组的值
      // fmt.Println( "数组numberArr的第", j ,"个值是", numberArr[ j ], "\n" )
  }

  //多维数组 和索引数组一样的 定长 和不定长
  var a = [][] int { {0,0,1}, {1,2}, {2,3,5,5,5,5,5},{1} }
  fmt.Println( a )


}
