package main
import (
	"fmt"
	"strings"
	"unicode"
)

const (
	easyAttem	= 8
	mediumAttem = 6
	hardAttem	= 4
)

func Interect() {
	fmt.Println("Welcome to Hangman")

	difficul := Difficul()
	attem := AttemDifficul(difficul)
	categories := Categories()
	category := Category(categories)
	word := RandWord(category)

	word = strings.ToUpper(word)
	ExpectedLetters := make([]bool, 26)
	GuessedLetters := make([]bool, 26)

	for _, ch := range word {
		if unicode.IsLetter(ch) {
			ExpectedLetters[ch-'A'] = true
		}
	}

	mistakes := 0
	for {
		DrawHangman(mistakes)
		drawWord(word, GuessedLetters)
		fmt.Print("\nAttemps left: %d\n", attem-mistakes)

		if mistakes >= attem {
			fmt.Println("\nYo lost! The word was:", word)
			break
		}
		
		if checkWin(ExpectedLetters, GuessedLetters) {
        	fmt.Println("\nYou won!") 
			break
    	}

		var input string
		fmt.Print("Enter a letter: ")
		fmt.Scanln(&input)
		if len(input) != 1 {
			fmt.Println("Please enter only one letter.")
			continue
		}

		letter := unicode.ToUpper(rune(input[0]))
		if letter < 'A' || letter > 'Z' {
			fmt.Println("Invalid input. Only letters are allowed.")
			continue
		}

		index := letter - 'A'
		if GuessedLetters[index] {
			fmt.Println("You already guessed this letter.")
			continue
		}

		GuessedLetters[index] = true
		if !ExpectedLetters[index] {
			mistakes++
			fmt.Println("Wrong guess!")
		} else {
			fmt.Println("Correct guess!")
		}
	}
}

func NoInterect(unknown, guessed string) {
	unknown = strings.ToUpper(unknown)
	guessed = strings.ToUpper(guessed)

	ExpectedLetters := make([]bool, 26)
	GuessedLetters := make([]bool, 26)

	for _, ch := range unknown {
		if unicode.IsLetter(ch) {
			ExpectedLetters[ch-'A'] = true
		}
	}
	for _, ch := range guessed {
		if unicode.IsLetter(ch) {
			GuessedLetters[ch-'A'] = true

		}
	}

	result := ""
	win := checkWin(ExpectedLetters, GuessedLetters)
	for _, ch := range unknown {
		if GuessedLetters[ch-'A'] {
			result += string(ch)
		} else {
			result += "*"
		}
	}

	status := "lose"
	if win {
		status = "win"
	}

	fmt.Printf("%s;%s\n", result, status)


}

func Difficul() string {
	var diff string
	for {
		fmt.Println("Choose difficulty (easy, medium, hard):")
		fmt.Scanln(&diff)

		if diff == "easy" || diff == "medium" || diff == "hard" {
			return diff
		}
		fmt.Println("Invalid difficulty, try again.")
	}
}

func AttemDifficul(diff string) int {
	switch diff {
	case "easy":
		return easyAttem
	case "medium":
		return mediumAttem
	case "hard":
		return hardAttem
	default:
		return mediumAttem
	}
}

func Category (categories []string) string {
	valid := make(map[string]bool)
	for _, c := range categories {
		valid[strings.ToLower(c)] = true

	}


	var c string
	for {
		fmt.Println("Choose category (or press Enter for random):")
		fmt.Scanln(&c)
		c = strings.ToLower(c)
		if c == "" || valid[c] {
			return c
		}
		fmt.Println("Invalid category, try again.")
	}
}

func checkWin(expected, guessed []bool) bool {
	for i := 0; i < 26; i++ {
		if expected[i] && !guessed[i] {
			return false
		}
	}
	return true
}

func drawWord(word string, guessed []bool) {
	for _, ch := range word {
		if guessed[ch-'A'] {
			fmt.Printf("%c ", ch)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}
