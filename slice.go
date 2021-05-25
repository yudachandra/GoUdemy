package main

import "fmt"

func main(){
  var months = [...]string{
    "Januari",
    "Februari",
    "Maret",
    "April",
    "Mei",
    "Juni",
    "Juli",
    "Agustus",
    "September",
    "Oktober",
    "November",
    "Desember",
  }

  var slice1 = months[:3]
  var slice2 = months[3:6]
  var slice3 = months[6:9]
  var slice4 = months[9:]

  fmt.Println(slice1)
    fmt.Println(slice2)
      fmt.Println(slice3)
        fmt.Println(slice4)

  var slice5 = append(slice4, "Bulan Baru")
  fmt.Println(slice5)

  newSlice := make([]string, 2, 5)

  newSlice[0] = "Yuda"
  newSlice[1] = "Chandra"

  fmt.Println(newSlice)
    fmt.Println(len(newSlice))
      fmt.Println(cap(newSlice))

  copySlice := make([]string, len(newSlice), cap(newSlice))
  copy(copySlice, newSlice)
  fmt.Println(copySlice)
}
