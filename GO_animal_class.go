package main

import(
  "fmt"
  "strings"
  "bufio"
  "os"
)

type Animal interface{
  Eat()
  Move()
  Speak()
}

type Cow struct{name string}
type Bird struct{name string}
type Snake struct{name string}

func (c Cow) Eat(){ fmt.Println("grass")}
func (b Bird) Eat(){ fmt.Println("worm")}
func (s Snake) Eat(){ fmt.Println("mice")}

func (c Cow) Move(){ fmt.Println("walk")}
func (b Bird) Move(){ fmt.Println("fly")}
func (s Snake) Move(){ fmt.Println("slither")}

func (c Cow) Speak(){ fmt.Println("moo")}
func (b Bird) Speak(){ fmt.Println("peep")}
func (s Snake) Speak(){ fmt.Println("hisss")}

func PrintInfo(a Animal, info string){
  switch info{
  case "eat":
    a.Eat()
  case "move":
    a.Move()
  default:
    a.Speak()
  }
}


func main()  {
  cows := []*Cow{}
  birds := []*Bird{}
  snakes := []*Snake{}

  fmt.Println("Please type 'newanimal <name of animal> <cow/bird/snake>' to create new animal")
  fmt.Println("(Please make sure that each animal that you create doesn't have the same name as animals you've created before)")
  fmt.Println("PLEASE CREATE ANIMAL BEFORE QUERYING, ALSO QUERY FOR ONLY ANIMALS WHICH YOU CREATED.")
  fmt.Println("or type 'query <name of animal> <eat/move/speak>' to find information")
  fmt.Println("(each word is separated by space)")
  fmt.Println("or type 'x' to exit")
  fmt.Println()
  for {
    fmt.Print("--RK--> ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()
    input = strings.ToLower(input)
    if input == "x" {break}

    inputs := strings.Fields(input)

    switch inputs[0] {
    case "newanimal":
      insert := "false"
      switch inputs[2] {
      case "cow":
        cow := new(Cow)
        cow.name = inputs[1]
        cows = append(cows,cow)
        insert = "true"
      case "bird":
        bird := new(Bird)
        bird.name = inputs[1]
        birds = append(birds,bird)
        insert = "true"
      default:
        snake := new(Snake)
        snake.name = inputs[1]
        snakes = append(snakes,snake)
        insert = "true"
      }
      if insert == "true"{
        fmt.Println("Created it!")
      }
    default:
      var a1 Animal
      for a := range cows {
        if cows[a].name == inputs[1]{
          a1 = cows[a]
        }
      }
      for a := range birds {
        if birds[a].name == inputs[1]{
          a1 = birds[a]
        }
      }
      for a := range snakes {
        if snakes[a].name == inputs[1]{
          a1 = snakes[a]
        }
      }
      if a1 == nil{
        fmt.Println("The animal you've searched is not in the list")
      }else{
        PrintInfo(a1,inputs[2])
      }
    }
    fmt.Println()
  }
}
