package main
import (
   "sync"
   "time"
   "math/rand"
   "fmt"
)

var counter int
var wg sync.WaitGroup

// Increment is a function to increase value of counter by 1 based on input val
func Increment(val int) {
	for i := 0; i < val*1000; i++ {
		counter++
	}
	fmt.Println(counter)
	wg.Done()
}
var wait sync.WaitGroup
var count int
func  increment(s string)  {
   for i :=0;i<10;i++ {
      x := count
      x++;
      time.Sleep(time.Duration(rand.Intn(4))*time.Millisecond)
      count = x;
      fmt.Println(s, i,"Count: ",count)

   }
   wait.Done()

}
// Explanation: We have a global variable counter that is incremented by one by two goroutines
// If we execute the program multiple times, it causes a race condition as there no
// way to know which Increment method is executed first and the results are always different
func main() {
	wg.Add(2)
	go Increment(100)
	go Increment(20)
	wg.Wait()

  wait.Add(2)
  go increment("foo: ")
  go increment("bar: ")
  wait.Wait()
  fmt.Println("last count value " ,count)
}
// package main
// import (
//     "fmt"
//     "sync" // to import sync later on
// )
// var GFG  = 0
//
// // This is the function we’ll run in every
// // goroutine. Note that a WaitGroup must
// // be passed to functions by pointer.
// func worker(wg *sync.WaitGroup, m *sync.Mutex) {
//     // Lock() the mutex to ensure
//     // exclusive access to the state,
//     // increment the value,
//     // Unlock() the mutex
//     m.Lock()
//     GFG = GFG + 1
//     m.Unlock()
//
//     // On return, notify the
//     // WaitGroup that we’re done.
//     wg.Done()
// }
// func main() {
//     // This WaitGroup is used to wait for
//     // all the goroutines launched here to finish.
//     var w sync.WaitGroup
//
//     // This mutex will synchronize access to state.
//     var m sync.Mutex
//
//     // Launch several goroutines and increment
//     // the WaitGroup counter for each
//     for i := 0; i < 1000; i++ {
//         w.Add(1)
//         go worker(&w, &m)
//     }
//     // Block until the WaitGroup counter
//     // goes back to 0; all the workers
//     // notified they’re done.
//     w.Wait()
//     fmt.Println("Value of x", GFG)
// }
