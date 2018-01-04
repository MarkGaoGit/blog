package main

import (
  "fmt"
)


func main(){
  a, b := 100, 200
  fmt.Println( "交换之前的值:", a ,"b的值是", b )

  swap ( &a, &b )
  fmt.Println( "交换值:", a ,"b的值是", b )

}

func swap ( x *int, y *int ){
  var temp int
  temp = *x   //保存X的值
  *x = *y
  *y = temp

}
