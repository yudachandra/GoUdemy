package main

import "fmt"

func main() {
  name := "Adi"

  if length := len(name); length <= 4 && name == "Yuda"{
    fmt.Println("Yuda")
  } else {
    fmt.Println("Nama Terlalu Panjang")
  }

  switch name {
  case "Yuda" : fmt.Println("Owner")
  case "Chandra" : fmt.Println("Owner Juga")
  default : fmt.Println("Bukan siapa-siapa")
  }

  // switch length := len(name); length <= 4 && name == "Yuda" {
  // case true : fmt.Println("Owner")
  // case false : fmt.Println("Bukan Owner")
  // }

  length := len(name);
  switch {
  case length < 5 : fmt.Println(name)
  case length > 5 : fmt.Println("Nama Terlalu Panjang")
  default : fmt.Println("Nama Sudah Benar")
  }
}
