package main

import (
  "fmt"
  "time"
)
func main(){
  present:= time.Now()
  fmt.Println("Today is -:",present.Format("Mon, Jan 1, 2006 at 01:29pm"))
  tmp_time:= time.Date(2019,time.August,24,11, 30, 55, 123456,time.Local)
  diff:= present.Sub(tmp_time)
  days:= int(diff.Hours()/24)
  fmt.Println("subtracted is ",diff,days)
  }
