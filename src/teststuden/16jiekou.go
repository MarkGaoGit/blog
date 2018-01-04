package main

import(
  "fmt"
)

type Phone interface{     //定义一个接口
  call();                 //  接口方法
}

type NokiaPhone struct{   //定义结构体
  phoneNumber string;     //变量
}


//接口实现方法
      //变量      结构名       方法名
func (nokiaPhone NokiaPhone) call(){
  var n NokiaPhone;
  n.phoneNumber = "18551857660";
  fmt.Println( "NOKIA call You" );    //方法
  fmt.Println( n.phoneNumber);    //方法
}

type Iphone struct{

}

func ( iphone Iphone) call(){
  fmt.Println( "Iphone Mobile" );
}

func main() {

  var phone Phone;
  phone = new( NokiaPhone );
  phone.call();

  phone = new( Iphone );
  phone.call();

}
