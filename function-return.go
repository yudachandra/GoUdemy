package main

import "fmt"

func main(){
  dompet := map[int]int{
    1:100000,
    2:50000,
    3:20000,
  }

  hargaBarang := 30000

  if hargaBarang > 50000{
    fmt.Println("Ada kembalian? ", jajan(hargaBarang, dompet[1]))
  } else {
    fmt.Println("Ada kembalian? ", jajan(hargaBarang, dompet[2]))
  }
}

func jajan (hargaBarang int, uangAwal int) bool{
  if hargaBarang < uangAwal {
    return true
  } else {
    return false
  }
}
