package main

// importing required packages
import (
   "sync"
   "time"
   "math/rand"
   "fmt"
)

// Defining variables
var wait sync.WaitGroup
var count int
var mutex sync.Mutex

//Function defination
func  increment(s string)  {
   for i :=0;i<10;i++ {
      mutex.Lock()
      x := count
      x++;
      time.Sleep(time.Duration(rand.Intn(20))*time.Millisecond)
      count = x;
      fmt.Println(s, i,"Count: ",count)
      mutex.Unlock()

   }
   wait.Done()

}
func main(){
   wait.Add(2)
   go increment("RISHI-: ")
   go increment("KEMNI-: ")
   wait.Wait()
   fmt.Println("Last count value " ,count)
}
