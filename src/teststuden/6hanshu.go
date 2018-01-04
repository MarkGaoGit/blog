package main

import (
  "fmt"
  "math"
)

// 函数的创建 和 调用
// 如果有默认值 需要设置其类型
// 如果函数不需要返回值就不需要设置返回类型

func orderCreate( sn string, money float64 ) (string, float64){
  return sn, money
}


func dx( num1, num2 int ) string {
  var res string
  if num1 > num2 {
    res = "num1 大"
  }else if num1 < num2 {
    res = "num2大"
  }else{
    res = "两个相等"
  }
  return res
}

// 函数传递值
//值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
func swap( x, y int ) int {
  var temp int
  temp = x
  x = y
  y = temp
  return temp
}

//函数引用传值 也就是使用 & 引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
func yycz( x *int, y *int) int {
  var temp int
  temp = *x           //这步只是存储 X的值 并不改变
  *x = *y             //x = y的值  = 2 把Y的值赋给X
  *y = temp           //y = x的值  = 1 把X的值赋给Y
  return temp
}



//函数当做值来使用 系统自带函数

func ns( x float64 ) float64{
   return math.Sqrt( x )
}


//闭包函数
func bbw() func() int {
   i := 0
   return func () int{
     i += 1
     return i
   }
}

func main(){
  //闭包函数
  bb := bbw()
  fmt.Println( bb() )
  fmt.Println( bb() )
  fmt.Println( bb() )
  fmt.Println( "`````````````````````````````````````````````````````````````````" )
  ac := bbw()
  fmt.Println( ac() )
  fmt.Println( ac() )
  fmt.Println( ac() )
  fmt.Println( ac() )

  sqrtNumber := ns( 4 )
  fmt.Println( sqrtNumber )

  swap1, swap2 := 1, 2
  var res1 = swap( swap1, swap2 );
  fmt.Print( res1, "\n" )

  fmt.Print( swap1, swap2, "\n" )


  yycz( &swap1, &swap2 )  //这个函数已经改变了 两个变量的值 因为加了& 所以 这里的也改变了

  fmt.Print( swap1, swap2, "\n" )



  sn, money := "2017102610082213256541", 33.5
  a, e := orderCreate( sn, money )    //传入多值 返回多值 如果可以择返回数组
  fmt.Println( a, e )

  var res string = dx( 11, 1)
  fmt.Println(res)

  var members string = "abcdefghigklmnopqrstuvwxyz"
  var Arr = [10] int {1,2,3,4,5,6}
  var lgs = len( members )
  var lgsArr = len( Arr )
  fmt.Println( lgs )
  fmt.Println( lgsArr )
}
