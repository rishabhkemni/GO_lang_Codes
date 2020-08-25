package main

import (
  "fmt"
  "errors"
  "math"
)


func find_sqrt(value float64)(float64, error) {
   if(value < 0){
      return 0, errors.New("Math: negative number passed to function find_sqrt")
   }
   return math.Sqrt(value), nil
}
func main() {
  var n float64
  for true{
    fmt.Println("Please enter a number to find square root. \n Type Ctrl+x to exit program")
    fmt.Scanln(&n)
    result, err:= find_sqrt(n)
    if err != nil {
       fmt.Println(err)
    } else {
       fmt.Println(result)
    }
  }
}
