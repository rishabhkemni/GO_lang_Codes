package main

import (
  "fmt"
)

func prod(v1 int,v2 int, c chan int){
  c <- v1*v2
}

func main(){
  c:= make(chan int)
  go prod(2,3,c)
  go prod(4,5,c)
  fmt.Println(c)
  a:= <-c
  b:= <-c
  fmt.Println(a,b)
  multi_num := a*b
  fmt.Println(multi_num)
}
