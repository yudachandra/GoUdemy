package main

import "fmt"

func main(){
  for i:=0; i < 2; i++{
   sayHello(i)
 }
}

func sayHello(jumlah int){
  fmt.Println("Halooo!!")

    for i:=0; i <= jumlah; i++{
      fmt.Println("Halooo", i)
    }
}
