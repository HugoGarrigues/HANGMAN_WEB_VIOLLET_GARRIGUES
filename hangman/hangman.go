package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Joueur struct {
	pseudo string
	mode   int
	essais int
}

func contains(slice []int, element int) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func MotAleatoire() string {
	fichier, err := os.Open("mot.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier !")
	}
	defer fichier.Close()

	scanner := bufio.NewScanner(fichier)
	var mots []string
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}

	rand.Seed(time.Now().UnixNano())
	return mots[rand.Intn(len(mots))]
}

func LettreEstPresente(lettre string, mot string) bool {
	return strings.Contains(mot, lettre)
}

func MotEstTrouve(motCache string) bool {
	return !strings.Contains(motCache, "_")
}

func masquerMot(mot string) string {
	n := len(mot)/2 - 1
	runes := []rune(mot)
	var indexes []int
	var revealed []rune
	for i := 0; i < n; i++ {
		index := rand.Intn(len(runes))
		for contains(indexes, index) {
			index = rand.Intn(len(runes))
		}
		indexes = append(indexes, index)
		revealed = append(revealed, runes[index])
	}
	for i := 0; i < len(runes); i++ {
		if !contains(indexes, i) {
			runes[i] = '_'
		}
	}

	return string(runes)
}
