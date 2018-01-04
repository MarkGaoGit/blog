package main

import(
  "fmt"
)

func main() {

  numbers := make([]int, 3, 5)
  printSlice( numbers )

  //空切片
  var nilSlice []int
  if nilSlice == nil{
    fmt.Println( "空切片" )
  }

  //切片截取
  messages := []int { 1,2,3,4,5,6,7,8,9,0 } //创建切片
  printSlice( messages )
  fmt.Println( "原始切片===", messages )
  fmt.Println( "打印子切片 索引1（包含） ~ 4（不包含）", messages[1:4] )
  fmt.Println( "messages[:3]==", messages[:3] )   //默认下限为0
  fmt.Println( "messages[4:]==", messages[4:] ) //默认上限为len

  messages0 := make([]int, 0, 5)
  printSlice( messages0 )

  messages1 := messages[:2] //打印子切片 0 ~２
  printSlice( messages1)

  messages2:= messages [2:5]
  printSlice( messages2)


}


//len 和 cap 函数
//len 获取长度  cap测量切片最长可达到多少
func printSlice ( x []int){
  fmt.Printf( "len=%d cap=%d slice=%v\n", len(x), cap(x), x )
}
