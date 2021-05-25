package main

import "fmt"

func main(){
  firstName, _ := getFirstName()
  //firstName, lastName := getFullName()
  fmt.Println(firstName)
}

func getFullName() (string, string) {
  return "Yuda", "Chandra"
}

func getFirstName() (string, string) {
  return "Yuda", "Chandra"
}
