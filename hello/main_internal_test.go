package main
import "testing"

func TestGreet_English(t *testing.T){
  want := "Hello World"
  got := greet("en")
  if got != want {
    t.Errorf("expected: %q, got: %q", want, got)
  }
}
func TestGreet_French(t *testing.T){
  want := "Bonjour"
  got := greet("fn")
  if got != want {
    t.Errorf("expected: %q, got: %q", want, got)
  }
}
