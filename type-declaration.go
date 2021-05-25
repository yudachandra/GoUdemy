package main

import "fmt"

func main(){
  type Tinggi int
  type Status bool

  var tinggiAduy Tinggi = 172
  var tinggiTyas Tinggi = 151
  var tinggiTotal int
  var statusAduy Status = true

  fmt.Println(tinggiAduy)
  fmt.Println(tinggiTyas)
  fmt.Println(statusAduy)

  //tinggiTotal = tinggiAduy + tinggiTyas

    var tinggiAduyBesok = 175
    var tinggiTyasBesok = 155
    tinggiTotal = tinggiAduyBesok + tinggiTyasBesok

  fmt.Println(tinggiTotal)
}
