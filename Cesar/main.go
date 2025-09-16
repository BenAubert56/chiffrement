package main

import "fmt"

// TD1 : Implémente un chiffrement de César en Go. Le chiffrement de César est une technique de chiffrement simple où chaque lettre dans le texte est décalée d'un nombre fixe de positions dans l'alphabet.

func main() {
	fmt.Println(caesar("abc", 3))   // def
	fmt.Println(deceasar("def", 3)) // abc
}

func caesar(word string, shift int) string {
	result := make([]rune, 0, len(word)) // Create a tab of runes to store the result
	for _, char := range word {          // For each character of the word
		if char >= 'a' && char <= 'z' { // Verify if the character is in lowercase
			shifted := ((int(char-'a') + shift) % 26) + 'a' // Store the modified caracter by recupering the place of the caractere and adding the shift. Moreover, a modulo 26 id doing to avoid circular shift and finally we add 'a' to get the ASCII value of the new character
			result = append(result, rune(shifted))          // Add the new character in the result
		} else if char >= 'A' && char <= 'Z' { //  Verify if the character is in uppercase
			shifted := ((int(char-'A') + shift) % 26) + 'A'
			result = append(result, rune(shifted))
		} else {
			result = append(result, char) // add the non modified character in the result
		}
	}
	return string(result)
}

func deceasar(word string, shift int) string {
	return caesar(word, 26-shift)
}
