package main

import "fmt"

func main(){
  a, b, c := getFullName()
  fmt.Println("First Name : ", a)
  fmt.Println("Last Name : ", b)
  fmt.Println("Age : "c)
}

func getFullName() (firstName string, lastName string, Age int){
  firstName = "Yuda"
  lastName = "Chandra"
  Age = 24
  return
}
