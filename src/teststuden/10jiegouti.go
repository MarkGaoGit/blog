package main

import (
  "fmt"
)

/*结构体 类似于PHP的类*/
type Prints struct{
  types string
  name string
  money float64
  year  int
}

//结构体作为函数参数
func printMessage ( p Prints ){
  fmt.Println( "传入函数", p.types )
  fmt.Println( "传入函数", p.name )
  fmt.Println( "传入函数", p.money )
  fmt.Println( "传入函数", p.year )
}

//结构体指针
func printMessagePtr ( p *Prints){
  fmt.Println( "使用指针传入", p.types )
  fmt.Println( "使用指针传入", p.name )
  fmt.Println( "使用指针传入", p.money )
  fmt.Println( "使用指针传入", p.year )
}


func main () {

  var Prints_Cancon Prints  //调用上面的类 也就是 new

  Prints_Cancon.types = "Cancon"
  Prints_Cancon.name = "Cancon E8901"
  Prints_Cancon.money = 10000.00
  Prints_Cancon.year = 2016

  fmt.Println( "重新赋值", Prints_Cancon.types )
  fmt.Println( "重新赋值", Prints_Cancon.name )
  fmt.Println( "重新赋值", Prints_Cancon.money )
  fmt.Println( "重新赋值", Prints_Cancon.year )

  //作为函数使用
  printMessage( Prints_Cancon )

  //使用指针传入

  printMessagePtr ( &Prints_Cancon )

}
