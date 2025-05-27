package main
import "fmt"

func main() {
  greeting := greet("fn")
  fmt.Println(greeting)
}

type language string

func greet(l language) string {
  switch l {
  case "en":
    return "Hello World"
  case "fn":
    return "Bonjour"
  default:
    return ""
  }
}
