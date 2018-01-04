package main

import(
  "fmt"
)

func main(){

  //1这种循环和PHP循环一样的
  length_number := 15
  for number := 0; number < length_number; number++ {
    if number == 13{
      continue;
    }
    fmt.Println( "getDocumentEmlentBuyId.type", number )
  }

  b := 15
  var a int
            // 长度从0 开始
  numbers := [10]int{1, 2, 3, 5, 2,34,}

  for a := 0; a < b; a++ {
    // fmt.Println(a);
  }


  for a < b {
    a++
    // fmt.Println(a)
  }

  for i, x := range numbers{
    if x > 2 {
      fmt.Println("第", i,"个值大于2", x)
    }else{
        fmt.Println("正常的值", x)
    }

  }

  var key, val int

  for key = 2; key <= 100; key++{

    for val = 2; val <= (key/val); val++{
        if ( key % val == 0 ){
          break;
        }
    }
    if( val > ( key / val ) ){
      // fmt.Println("素数\n",key);
    }

  }

  var M, m int
  M = 1

  A: for M < 100 {
      M++
      for m =2 ; m < M ; m++{
          if M%m == 0 {
              goto A
            }
      }
      fmt.Println( "素数", M )

  }

  //for 改为 while循环
  sumNumber := 1;
  for sumNumber < 1000{
	sumNumber += sumNumber
  }
  fmt.Println(sumNumber)


}
