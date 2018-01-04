package main

import(
  "fmt"

)



func main(){
  var a int = 10
  var b int = 20
  var c,d,e,f,g,h,i, j, k, l, m, n,o,p int
  c = a + b //30
  d = b - a //10
  e = b * a //200
  f = e / a //20
  g = f % 3 //2
  g++
  f--
  fmt.Println( c, d, e, f, g,h,i );

  k = 10
  l = 100
  m = 100
  n = 100
  o = 100
  p = 100

  h = e   //200
  i += d  //10
  j -= h  //-200
  k *= a  //100
  l /= e  //0
  m %= g  //1
  // n <<= a
  // o >>= a
  p &= a

  fmt.Println( h,i, j,k,l,m,n,o,p )

//=== and !=  >   <   >=   <=
  if a == b {
    fmt.Println("==")
  }else if a != 10 {
    fmt.Println("!=")
  }else if a < 9 {
    fmt.Println("<")
  }else if a > 20 {
    fmt.Println(">")
  }else if b >= 21{
    fmt.Println(">=")
  }else if b <= 20{
    fmt.Println("<=")
  }

  var type1 bool = true
  var type2 bool = false

  if type1 && type2{
    fmt.Println("都是true")
  }else if type1 || type2 {
    fmt.Println("1个是true就可以了")
  }

  if !type2 {
    fmt.Println("!Not")
  }

  var aa int = 4
  var bb int32
  var cc float32
  var ptr *int

  /* 运算符实例 */
  fmt.Printf("第 1 行 - a 变量类型为 = %T\n", aa );
  fmt.Printf("第 2 行 - b 变量类型为 = %T\n", bb );
  fmt.Printf("第 3 行 - c 变量类型为 = %T\n", cc);

  /*  & 和 * 运算符实例 */
  ptr = &aa    /* 'ptr' 包含了 'a' 变量的地址 */

  fmt.Printf("aa 的值为  %d\n", aa);
aa = 100
  fmt.Printf("*ptr 为 %d\n", *ptr);

}
