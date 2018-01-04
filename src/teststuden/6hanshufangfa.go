package main

import (
  "fmt"
)

/*定义一个函数*/
// 类似于PHP的 $this->xxx = ''
type Circle struct {
  radius float64
}

func main(){

  var c1 Circle       //类似于PHP的new Class()
  c1.radius = 10.00   //类似于PHP的$this->xxx = 10
  fmt.Println( "Area of Circle(C1)", c1.getArea() )

}

// 该方法属于 Circle类型对象中的方法
//    name daat_type
func ( c Circle ) getArea() float64 {       //类似于 PHP的继承  类名 方法名  return结果类型

  return 3.14 * c.radius * c.radius

}
