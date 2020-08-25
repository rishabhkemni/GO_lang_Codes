package main

import (
  "math"
  "fmt"
)

func GenDisplaceFn(acceleration, ini_vel, ini_disp float64) func (float64) float64{
  fn:= func(time float64) float64{
    return 0.5 * acceleration * math.Pow(time,2) + ini_vel * time + ini_disp
  }
  return fn
}

func main(){
  var (
    acceleration float64
    initial_velocity float64
    initial_displacement float64
    time float64
  )
  fmt.Println("Enter acceleration")
  fmt.Scanln(&acceleration)
  fmt.Println("Enter initial_velocity")
  fmt.Scanln(&initial_velocity)
  fmt.Println("Enter initial_displacement")
  fmt.Scanln(&initial_displacement)
  tmp_func:= GenDisplaceFn(acceleration,initial_velocity,initial_displacement)
  fmt.Println("Enter time for which displacement needed to be calculated")
  fmt.Scanln(&time)
  displacement :=tmp_func(time)
  fmt.Println("Final displacement is ",displacement)
}
