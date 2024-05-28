package main

import (
	"fmt"
	"log"
	"regexp"
	"unicode/utf8"
)

// Чтение слова из консоли
// прроверка на существование введенного слова
// проверка попытки и вывод результата

// СПИСОК СЛУЧАЕВ
// 1. нет входящих букв
// 2. слово не на своем месте
// 3. слово на правильном месте
// 4. Проигрыш и вывод правильного слова
// 5. выигрыш
// 6. Длина больше или меньше 5 //сделано
// 7. Проверка на язык //сделано

func checkLen(word string) bool {
	// fmt.Println(len(word)) // считаем байты
	// fmt.Println(utf8.RuneCountInString(word)) // считаем количество символов
	if utf8.RuneCountInString(word) != 5 {
		fmt.Printf("В слове должно быть 5 букв")
		return false
	}
	return true
}

func checkLanguage(word string) bool {
	var letters []string
	var isRussian = true
	for _, letter := range word {
		letters = append(letters, string(letter))
	}

	for _, rune := range letters {
		re := regexp.MustCompile("[А-Яа-я]")
		if !re.MatchString(rune) {
			isRussian = false
			fmt.Println("Слово должно быть русским языком написано")
			break
		}
	}
	return isRussian
}

func checkTry(word string, wordx string) []string {
	var letters []string
	var lettersx []string
	var result []string

	for _, letter := range word {
		letters = append(letters, string(letter))
	}
	log.Println(letters)
	log.Println(lettersx)
	for _, letterx := range wordx {
		lettersx = append(lettersx, string(letterx))
	}
	result = make([]string, len(letters))
	usedInWordx := make([]bool, len(wordx))

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



func main() {
	wordx := "вссср"
	for try := 1; try <= 6; try++ {
		fmt.Println("Введите 5-буквенное слово на русском языке")
		var word string
		fmt.Scanf("%s\n", &word)
		if !checkLen(word) || !checkLanguage(word) {
			break
		}
		checkTry(word, wordx)
	}
}
