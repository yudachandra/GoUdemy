package main

import "fmt"

func main() {
  person := map[string]string{
    "name" : "Yuda",
    "address" : "Chandra",
  }

  fmt.Println(person)
    fmt.Println(person["name"])
      fmt.Println(person["address"])

  person["title"] = "Programmer"

    fmt.Println(person)
    fmt.Println(len(person))

    delete(person, "title")

    fmt.Println(person)
    fmt.Println(len(person))
}
