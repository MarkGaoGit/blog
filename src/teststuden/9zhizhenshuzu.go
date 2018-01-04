package main

import (
  "fmt"
)

const MAX int = 3

func main(){
  arr := []int{1,2,3}
  var i int
  var ptr [MAX] *int

  for i = 0; i < MAX; i++ {
    ptr[i] = &arr[i]
  }

  for i = 0; i < MAX; i++{
    fmt.Println( "指针数组ptr的第", i,"个值是", *ptr[i] )
  }


  //这样直接使用 指针数组不是更方便吗
  a := &arr
  b := *a
  

  fmt.Println( b )

}
