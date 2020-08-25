package main

import (
  "fmt"
)
type vehicle interface {
  accelerate()
}

func test_1(v vehicle){
  fmt.Println(v)
}

type car struct {
  model string
  color string
}

func (c car) accelerate(){
  fmt.Println("Accelerating?")

}

type toyota struct {
  model string
  color string
  speed int
}

func (t toyota) accelerate(){
  fmt.Println("Hi! I am toyota, I accelerate faster than fast :)")
}

func main(){
  c1:=car{"ferrari","red"}
  t1:=toyota{"Toyota","Red",100}
  c1.accelerate()
  t1.accelerate()
  test_1(c1)
  test_1(t1)
}
