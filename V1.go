package main

import (
  "fmt"
  "strings"
)

func main() {
  var input string
 input = "ciao a tutti"
  //utiliziamo una funzione per fare l'acronimo
  acronimo := genera_acronimo(input)
  fmt.Println("L'acronimo generato è:", acronimo)
}

func genera_acronimo(input string) string {
  //splittiamo le parole
  parole := strings.Split(input, " ")
  var acronimo string
  for _, parola := range parole {
    acronimo += string(parola[0])
  }
  return strings.ToUpper(acronimo)
}

