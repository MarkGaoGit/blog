package main

import(
  "fmt"
)

func main(){
  // a := 10
  // fmt.Println( &a )

  //指针通常缩写为ptr

  var a int = 20    //实际变量
  var ip *int       //指针变量

  ip = &a           //指针变量的存储地址

  fmt.Printf( " a变量的地址是：%x\n ", &a )    //输出a的内存地址
  fmt.Printf( "ip变量存储的指针地址：%x\n", ip )  //ip = a的内存地址

  fmt.Printf( "*ip的变量值:%d\n", *ip )     //指针指向a的内存地址 储存的值 也就是20

  var ptr *int


  // ptr 为指针    *ptr为指针所指向的值
  if ptr == nil{
      fmt.Println( " 空指针打印 ", ptr )
  }




}
