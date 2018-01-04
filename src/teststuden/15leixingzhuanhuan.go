package main

import(
  "fmt"
)

func main() {

  //转换格式
  // 类型       表达式
  // type_name ( expression )

  var sum int = 17;
  var count int = 5;
  var mean float32;

  mean = float32( sum ) / float32( count );     //类型转换

  // mean = sum / count  //不转类型会报错
  fmt.Println( mean );

}
