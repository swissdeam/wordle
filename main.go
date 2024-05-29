package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

// Проверка длины слова
func checkLen(word string) bool {
	if utf8.RuneCountInString(word) != 5 {
		fmt.Println("В слове должно быть 5 букв")
		return false
	}
	return true
}

// Проверка языка слова
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
		// log.Println(result, "промежуточный")
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

func loadWords(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if utf8.RuneCountInString(word) == 5 {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func getRandomWord(words []string) string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	return words[rng.Intn(len(words))]
}

func isWordInList(word string, words []string) bool {
	for _, w := range words {
		if word == w {
			return true
		}
	}
	return false
}

func main() {
	words, err := loadWords("russian2.txt")
	if err != nil {
		log.Fatalf("Ошибка при загрузке слов: %v", err)
	}
	wordx := getRandomWord(words)
	fmt.Println("Загаданное слово:", wordx)

	for try := 1; try <= 6; try++ {
		fmt.Println("Введите 5-буквенное слово на русском языке")
		var word string
		fmt.Scanf("%s\n", &word)
		word = strings.ToLower(word) // Нижний регистр

		if !checkLen(word) || !checkLanguage(word) || !isWordInList(word, words) {
			fmt.Println("Попробуйте снова.")
			try--
			continue
		}

		result := checkTry(word, wordx)
		printResult(word, result)

		//выигрыш
		if strings.Join(result, "") == "🟢🟢🟢🟢🟢" {
			fmt.Println("Поздравляем! Вы угадали слово!")
			return
		}
	}

	fmt.Printf("Игра окончена. Вы не угадали слово. Загаданное слово: %s\n", wordx)
}
