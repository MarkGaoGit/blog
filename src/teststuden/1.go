package main

import (
  "fmt"
)

func init(){
  fmt.Println( "Init Function!" );
}

func main(){
  var age int = 100;
  fmt.Println( age );

  var is_open bool = true;
  if( is_open ){
    fmt.Println("Open");
  }else{
    fmt.Println("Close");
  }
  var money float32 = 34.5;
  fmt.Println(money);

  var a = "一二三";
  var b = "abc123";
  var c = 10;
  fmt.Println(a);
  fmt.Println(b);
  fmt.Println(c);
  d, e, f := "Mark", "的", 123;
  fmt.Println(d)
  fmt.Println(e)
  fmt.Println(f)


  is_open = true;
  var provite string;
  if( is_open ){
    provite = "Henan!";
  }else{
    provite = "Beijing!";
  }

  fmt.Println(provite);


  integral := 100;
  // level := "E会员";

  fmt.Println(integral);

  hotel_name := `淮安`;
  fmt.Println(hotel_name);

}
