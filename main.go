package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode/utf8"
)

// Проверка длины слова
func checkLen(word string) bool {
	if utf8.RuneCountInString(word) != 5 {
		fmt.Printf("В слове должно быть 5 букв\n")
		return false
	}
	return true
}

// Проверка языка слова (должно быть написано кириллицей)
func checkLanguage(word string) bool {
	re := regexp.MustCompile("^[А-Яа-я]+$")
	if !re.MatchString(word) {
		fmt.Println("Слово должно быть написано русскими буквами")
		return false
	}
	return true
}

// Проверка попытки и формирование результата
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

// Печать результата
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
		word = strings.ToLower(word) // Приводим слово к нижнему регистру

		// Проверка длины и языка, даем возможность повторного ввода
		if !checkLen(word) || !checkLanguage(word) {
			fmt.Println("Попробуйте снова.")
			try-- // Не засчитываем неудачную попытку
			continue
		}

		result := checkTry(word, wordx)
		printResult(word, result)

		// Проверка на выигрыш
		if strings.Join(result, "") == "🟢🟢🟢🟢🟢" {
			fmt.Println("Поздравляем! Вы угадали слово!")
			return
		}
	}

	fmt.Printf("Игра окончена. Вы не угадали слово. Загаданное слово: %s\n", wordx)
}
