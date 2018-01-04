package main

import(
  "fmt"
)

func main() {

  //声明格式         key     val    默认的map 是nil 空的
  var mapName map [string] string
                  //    key_type value_type
  map_name := make( map[string] string)   //第二种命名

  //nil map不可以用来存放键值对
  fmt.Println( mapName )
  fmt.Println( map_name )

  // var printName map[string] string
  var printName = make( map[string] string )

  printName["Hp"] = "E901"
  printName["Dell"] = "Dell-H3000"
  printName["Aus"] = "A584"
  printName["Cancon"] = "dgcz999"
  printName["Acer"] = "AGe450"

  //使用 k 输出 val 值
  for v := range printName{
    fmt.Println( v, printName[v] )
  }

  //判断值是否存在这个集合中
  s5, ok := printName["Acer"]

  if (ok ){
    fmt.Println("存在", s5)
  }else{
    fmt.Print("不存在")
  }

  fmt.Println( printName )


  //delete 删除数组的元素
  delete( printName, "Hp" )
  fmt.Println( printName )


}
