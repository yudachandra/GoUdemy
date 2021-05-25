package main

import "fmt"
import "strconv"

func main() {
  counter := 0

  for counter <= 5 {
    fmt.Println("Line :" + strconv.Itoa(counter))
    counter++
  }

  for counter2 := 0; counter2 <= 5; counter2++ {
    fmt.Println("Line :" + strconv.Itoa(counter2))
  }

  slice := []string{"tingky wingky", "dipsy", "lala", "po"}

  // for i := 0; i < len(slice); i++ {
  //   fmt.Println(slice[i])
  // }

  for j, values := range slice{
    fmt.Println("Index", j, "=", values)
  }

  for _, values := range slice{
    fmt.Println(values)
  }

  teletubbies := map[string]string{
    "Teletubies":"Tingky Wingky",
    "Age":"10",
    "Gender":"Male",
  }
  for i, value := range teletubbies {
    fmt.Println(i, "nya", value)
  }

  person := make(map[string]string)
  person["Teletubies"] = "Tingky Wingky"
  person["Age"] = "10"
  person["Gender"] = "Male"

  for key, value := range person {
    fmt.Println(key, "=", value)
  }
}
