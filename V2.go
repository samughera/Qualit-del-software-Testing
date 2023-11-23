package main

import (
	"fmt"
	"strings"
)

func main() {
	// Chiedi all'utente di inserire una frase
	fmt.Print("Inserisci una frase: ")
	var input string
	fmt.Scanln(&input)

	// Utilizziamo una funzione per fare l'acronimo
	acronimo := generaAcronimo(input)
	fmt.Println("L'acronimo generato è:", acronimo)
}

func generaAcronimo(input string) string {
	// Splittiamo le parole
	parole := strings.Split(input, " ")
	var acronimo string
	for _, parola := range parole {
		acronimo += string(parola[0])
	}
	return strings.ToUpper(acronimo)
}
