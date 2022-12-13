package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func contains(slice []int, element int) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func Verif_lettre(lettre string) bool {
	if lettre >= "a" && lettre <= "z" || lettre >= "A" && lettre <= "Z" {
		return true
	}
	return false
}

func MettreEnMajuscule(lettre string) string {
	if lettre >= "a" && lettre <= "z" {
		lettre = strings.ToUpper(lettre)
	}
	return lettre
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

func MasquerMot(mot string) string {
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

func AfficheMotAvecLettreTrouvee(lettre string, mot string, motCache string) string {
	var motCacheTemporaire string
	for i := 0; i < len(mot); i++ {
		if string(mot[i]) == lettre {
			motCacheTemporaire += lettre
		} else {
			motCacheTemporaire += string(motCache[i])
		}
	}
	return motCacheTemporaire
}

func LancementDuJeu(essais int) {
	var lettre string
	mot := MotAleatoire()
	nouveaumot := MasquerMot(mot)
	for essais < 0 || !MotEstTrouve(nouveaumot) {
		fmt.Println(nouveaumot)
		fmt.Scan(&lettre)
		if Verif_lettre(lettre) {
			lettre = MettreEnMajuscule(lettre)
			if LettreEstPresente(lettre, mot) {
				nouveaumot = AfficheMotAvecLettreTrouvee(lettre, mot, nouveaumot)
			} else {
				essais--
			}
		}
	}
}
