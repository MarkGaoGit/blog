package main

import(
  "fmt"
)

func main() {

  a := make(chan bool);

  go func(){
      fmt.Println("func has been calld")
      close(a)
  }()
  <-a
  fmt.Println("Ok Success")

}
