package main

import (
  "fmt"
  "unsafe"
)

//枚举类型定义常量
const(
  a = "abc"
  b = len(a)
  c = unsafe.Sizeof(a)
)

const(
  d = iota    //0
  e           //1
  f           //2
  g           //3
  h = "ha"    //ha    重新赋值
  i           //ha
  j           //ha
  k = 100     //100  重新赋值
  l           //100
  m           //100
  n = iota    //0   //出现iota回复之前的 +1 累加
  o           //1
)

const(
  p = 1<< iota  // p = 1 << 0   位运算
  q = 3<< iota  // q = 3 << 1
  r
  s
)

func main(){
  const prv_name = "上海"
  const class_name, class_long, class_type = "类名", "100byt", "interface"
  const DEFAULTSEX = 1
  const WIDTH, HEIGHT, LENGTH = 3,5,11
  volume := WIDTH * HEIGHT * LENGTH

  fmt.Println( prv_name, class_name, class_long, class_type )
  fmt.Println( DEFAULTSEX )
  fmt.Println( "体积为:", volume )
  fmt.Println( a, b, c )
  fmt.Println( d, e, f, g, h, i, j, k, l, m, n, o )
  fmt.Println( p, q, r, s )
}
