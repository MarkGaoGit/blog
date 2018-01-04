package main

import(
  "fmt"
)

func main() {
  var nums = []int { 1,2,3,4,5}
  sum := 0

  for _, num := range nums{     //_ 空白符号省略了 index
    sum += num
  }

  fmt.Println( "range求和nums:", sum )

  for i, num := range nums{
    if num == 4{
      fmt.Println( "index:", i )
    }
  }


  //range 对用在map集合的键值对上
  kvs := map[string]string { "a": "apple", "b": "banana" }
  for k, v := range kvs{
    fmt.Printf( " %s -> %s\n ", k, v )
  }

  //range用在枚举Unicode类型 第一个是字符的索引 第二个是字符本身

  for k, v := range "hello world"{
    fmt.Println( k,v )
  }

}
