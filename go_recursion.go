package main
import (
   "fmt"
)
func main() {
  var n = 0
  fmt.Print("Enter number to find factorial. \n")
  fmt.Scanln(&n)
   fmt.Println("Factorial is -: ",factorial(n))
}
func factorial(num int ) int{
   if num == 0{
      return 1
   }
   return num*factorial(num-1)
}
