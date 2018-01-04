package main

import (
  "fmt"
)

func main(){

  var a int
  var ptr *int
  var pptr **int

  a = 100
  ptr = &a
  pptr = &ptr

  fmt.Println( "默认值a", a )
  fmt.Println( "指针ptr" , *ptr )
  fmt.Println( "指针 指向 指针变量的值", **pptr )



}
