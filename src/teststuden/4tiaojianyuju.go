package main

import(
  "fmt"
)

var a, b = true, false
var c, d = 30, 20

func main(){
  if a {
    fmt.Println("a是true")
  }else{
    fmt.Println("a是false")
  }

  if c < d{
    fmt.Println("c < d")
  }else{
    fmt.Println("c > d")
  }


  var grade string = "B"
  var marks int = 70

  switch marks {
     case 90: grade = "A"
     case 80: grade = "B"
     case 50,60,70 : grade = "C"
     default: grade = "D"
  }

  switch {
     case grade == "A" :
        fmt.Printf("优秀!\n" )
     case grade == "B", grade == "C" :
        fmt.Printf("良好\n" )
     case grade == "D" :
        fmt.Printf("及格\n" )
     case grade == "F":
        fmt.Printf("不及格\n" )
     default:
        fmt.Printf("差\n" );
  }
  fmt.Printf("你的等级是 %s\n", grade );

//witch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
  var x interface{}
  switch i := x.(type) {
  case nil :
    fmt.Printf("类型",i)
  case float32:
    fmt.Printf("类型float32")

  }

//select 循环语句没看懂 其中涉及 chan 

}
