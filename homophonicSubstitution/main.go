package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strings"
	"unicode"
)

// Dictionnaire homophonique: chaque lettre -> plusieurs codes possibles.
var homo = map[rune][]string{
	'A': {"10", "57"},
	'B': {"11"},
	'C': {"12", "58"},
	'D': {"13"},
	'E': {"20", "21", "22", "23"}, // E très fréquent -> plus de codes
	'F': {"14"},
	'G': {"15"},
	'H': {"16"},
	'I': {"24", "25"},
	'J': {"17"},
	'K': {"18"},
	'L': {"26"},
	'M': {"27"},
	'N': {"28"},
	'O': {"29", "30"},
	'P': {"31"},
	'Q': {"32"},
	'R': {"33"},
	'S': {"34", "59"},
	'T': {"35"},
	'U': {"36"},
	'V': {"37"},
	'W': {"38"},
	'X': {"39"},
	'Y': {"40"},
	'Z': {"41"},
}

// Inverse la map afin de trouver à partir d'un code la lettre
var inv = map[string]rune{}

// Executé automatiquement en go avant tout :
// Check si pas de doublons dans les codes.
// Ex : si Z = 41 et E = 41 ça log une erreur
func init() {
	for r, codes := range homo {
		for _, c := range codes {
			if _, ok := inv[c]; ok {
				log.Fatalf("code %s dupliqué", c)
			}
			inv[c] = r
		}
	}
}

// tirage uniforme dans [0, n)
func randint(n int) int {
	if n <= 0 {
		return 0
	}
	bn, err := rand.Int(rand.Reader, big.NewInt(int64(n))) // Génère un nombre aléatoire (big int)
	if err != nil {
		return 0
	}
	return int(bn.Int64()) // return le int du big int
}

// Encode homophonique.
// - Lettres -> code numérique (choix aléatoire parmi la liste).
// - Espace -> "_".
// - Autre caractère non lettre -> jeton littéral tel quel.
func EncodeHomophonic(plain string) string {
	var tokens []string
	for _, r := range plain {
		switch {
		case r == ' ':
			tokens = append(tokens, "_")
		case unicode.IsLetter(r):
			u := unicode.ToUpper(r)
			codes, ok := homo[u]
			if !ok {
				// lettre non mappée -> littéral
				tokens = append(tokens, string(r))
				continue
			}
			tokens = append(tokens, codes[randint(len(codes))])
		default:
			// ponctuation / chiffres -> littéral
			tokens = append(tokens, string(r))
		}
	}
	return strings.Join(tokens, " ")
}

// Decode homophonique (selon règles ci-dessus).
func DecodeHomophonic(cipher string) string {
	var b strings.Builder
	for _, tok := range strings.Fields(cipher) {
		if tok == "_" {
			b.WriteRune(' ')
			continue
		}
		// si tok est un code numérique connu -> lettre
		if r, ok := inv[tok]; ok {
			b.WriteRune(r)
			continue
		}
		// sinon, tok est un littéral (ponctuation, etc.)
		b.WriteString(tok)
	}
	return b.String()
}

func main() {
	plain := "Ceci est, un test!"
	ciph := EncodeHomophonic(plain)
	deco := DecodeHomophonic(ciph)

	fmt.Println("PLAIN:", plain)
	fmt.Println("CIPH :", ciph)
	fmt.Println("DECO :", deco)
}
