package main

import (
	"fmt"
	"unicode"
)

// Fonction cesar réutilisée pour le chiffrement de Vigenère
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

// calcule les décalages de la clé
func keyShifts(key string) []int {
	shifts := make([]int, 0, len(key))
	for _, char := range key {
		if char >= 'a' && char <= 'z' {
			shifts = append(shifts, int(char-'a'))
		} else if char >= 'A' && char <= 'Z' {
			shifts = append(shifts, int(char-'A'))
		}
	}
	return shifts
}

// Vigenère chiffrement
func vigenereEncrypt(plain, key string) string {
	shifts := keyShifts(key)
	result := make([]rune, 0, len(plain))

	j := 0
	for _, char := range plain {
		if unicode.IsLetter(char) {
			shift := shifts[j%len(shifts)]
			// applique César sur une lettre
			enc := []rune(caesar(string(char), shift))
			result = append(result, enc[0])
			j++
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}

// Vigenère déchiffrement
func vigenereDecrypt(cipher, key string) string {
	shifts := keyShifts(key)
	result := make([]rune, 0, len(cipher))

	j := 0
	for _, char := range cipher {
		if unicode.IsLetter(char) {
			shift := (26 - shifts[j%len(shifts)]) % 26
			dec := []rune(caesar(string(char), shift))
			result = append(result, dec[0])
			j++
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}

func main() {
	plain := "Hello world"
	key := "KEY"

	ciph := vigenereEncrypt(plain, key)
	deco := vigenereDecrypt(ciph, key)

	fmt.Println("PLAIN:", plain)
	fmt.Println("CIPH :", ciph)
	fmt.Println("DECO :", deco)
}
