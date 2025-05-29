package main
import "testing"

func TestGreet(t *testing.T) {
  type testCase struct{
    lang language
    want string
  }
  var tests = map[string]testCase{
    "English": {
      lang: "en",
      want: "Hello world",
    },
    "French": {
      lang: "fn",
      want: "Bonjour",
    },
    "Spanish": {
      lang: "el",
      want: "Hola",
    },
    "Empty": {
      lang: "",
      want: `unsupported language: ""`,
    },
  }
  for name, tc := range tests {
    t.Run(name, func(t *testing.T) {
      got := greet(tc.lang)
      if got != tc.want {
        t.Errorf("expected: %q, got: %q", tc.want, got)
      }
    })
  }

}
