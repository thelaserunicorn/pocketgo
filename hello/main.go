package main
import (
  "fmt"
  "flag"
)


func main() {
  var lang string
  flag.StringVar(&lang, "lang", "en", "choose from language: en, fn, el")
  flag.Parse()
  greeting := greet(language(lang))
  fmt.Println(greeting)
}

type language string

var phrasebook = map[language]string {
  "el": "Hola",
  "en": "Hello world",
  "fn": "Bonjour",
}

func greet(l language) string {
  greeting, ok := phrasebook[l]
  if !ok {
    return fmt.Sprintf("unsupported language: %q", l)
  }
  return greeting
}
