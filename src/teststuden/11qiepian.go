package main

import(
  "fmt"
)

// make([]T, length, [capacity]) 切片定义格式 T:数组类型  length数组长度也是默认的切片长度 capacity：指定容量 可选


func main() {
  var qiepian []int   //定义一个切片 未指定大小
  
  var slice []int = make([]int, 5)  //使用make创建切片 并指定初始大小

  slices := make( []string, 3, capacity ) // 指定容量， 第三个参数可选
  slice := make( []int, 5)    //简便的创建切片 和上面的结果一样

.

  s := [] int { 1, 2, 3, 4}   //初始化切片 cap = len = 4

  a := arr[:] //初始化 切片a 是数组arr的引用

  b := arr[startIndex : endIndex] //数组 0~ 最后一个 的元素创建新的切片

  c := arr[startIndex:] //缺省endIndex时表示一直到arr的最后一个元素

  d := arr[:endIndex] //缺省startIndex时表示从arr的第一个元素开始

  e := s[startIndex : endIndex]   //通过切片s初始化e

  s := make([]int, len , cap)

  fmt.Println(s)
  fmt.Println(b)

}
