package main

import(
  "fmt"
)

//切片截取函数 append(追加)  和  copy(复制)
func main() {
  var number[]int
  printFunc( number )

  number = append( number, 0 )   //允许追加空切片
  printFunc( number )

  number = append( number, 1 )   //添加一个元素
  printFunc( number )

  number = append( number, 2,3,4,5,6 )   //向切片添加多个元素
  printFunc( number )

  numbers := make([]int, len( number ) * 2, ( cap( number ) ) * 2  )    //创建切片numbers 长度 和 容量 均是number切片的2倍
  copy( numbers, number )     //把number切片的内容 复制给 numbers切片
  printFunc( numbers )
}


func printFunc( x[]int ){

  fmt.Printf( " len=%d cap= %d slice=%v\n ", len(x), cap(x), x )

}
