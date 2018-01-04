package main

import (
  "fmt"
)

/*外部使用var声明的变量为全局变量*/
var x, y int;
var(    //因式分解法 一般用户生命全局变量
  a int;
  b bool
)

var c, d int = 123, 444;

func main(){

  e, f := 555, "Test";    //只可以在函数内部使用这种命名方法

  fmt.Println(x,y,a,b,c,d,e,f);
}
