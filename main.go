package main

import (
	"fmt"
	"log"
	"regexp"
	"unicode/utf8"
)

func checkLen(word string) bool {
	if utf8.RuneCountInString(word) != 5 {
		fmt.Println("В слове должно быть 5 букв")
		return false
	}
	return true
}

func checkLanguage(word string) bool {
	re := regexp.MustCompile("[А-Яа-я]")
	for _, letter := range word {
		if !re.MatchString(string(letter)) {
			fmt.Println("Слово должно быть написано русскими буквами")
			return false
		}
	}
	return true
}

func checkTry(word string, wordx string) []string {
	letters := []rune(word)
	lettersx := []rune(wordx)
	result := make([]string, len(letters))
	usedInWordx := make([]bool, len(lettersx))

	for i := 0; i < len(letters); i++ {
		if letters[i] == lettersx[i] {
			result[i] = "🟢"
			usedInWordx[i] = true
		}
	}

	for i := 0; i < len(letters); i++ {
		if result[i] == "🟢" {
			continue
		}

		matched := false
		for j := 0; j < len(lettersx); j++ {
			if letters[i] == lettersx[j] && !usedInWordx[j] {
				result[i] = "🟡"
				usedInWordx[j] = true
				matched = true
				break
			}
		}
		if !matched {
			result[i] = "⚫"
		}
		log.Println(result, "промежуточный")
	}

	return result
}

func printResult(word string, result []string) {
	for _, letter := range word {
		fmt.Printf("%c  ", letter)
	}
	fmt.Println()
	for _, r := range result {
		fmt.Printf("%s ", r)
	}
	fmt.Println()
}

func main() {
	wordx := "вссср"
	for try := 1; try <= 6; try++ {
		fmt.Println("Введите 5-буквенное слово на русском языке")
		var word string
		fmt.Scanf("%s\n", &word)
		if !checkLen(word) || !checkLanguage(word) {
			break
		}
		result := checkTry(word, wordx)
		printResult(word, result)
	}
}
