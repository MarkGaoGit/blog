package main

import (
  "fmt"
  "math"
)


//全局变量
const(
  name = "张三"
  sex = 9
)

func sqrtNumber( x float64) float64 {
  return math.Sqrt( sex )
}

// 不同类型的局部和全局变量默认值为： int 0   float32 0   pointer nil 

// 局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑
func prov ( name string ) string{
  var provname , countrys string = "大陆", "台湾"     //局部变量
  var sex = 4 //这里的sex将会比全局的sex优先级高
  if name == "上海"{
    return provname
  }else if name == "高雄"{
    return countrys
  }else{
    return "海外"
  }


}


func main(){

  //全局变量
  sqrtRes := sqrtNumber( sex )
  fmt.Println( sqrtRes )

  //函数形参
  provFunc := prov( "高雄" )
  fmt.Println( provFunc )

}
