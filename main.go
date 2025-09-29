package game

import (
	"fmt"
	"strings"
	"unicode"
)


type GameState struct {
	Word           string   
	GuessedLetters []rune   
	WrongAttempts  int      
	MaxAttempts    int     
	GameOver       bool    
	Won            bool     
}


func NewGame(word string, maxAttempts int) *GameState {
	return &GameState{
		Word:           strings.ToUpper(word),  
		GuessedLetters: make([]rune, 0),        
		WrongAttempts:  0,                      
		MaxAttempts:    maxAttempts,            
		GameOver:       false,                 
		Won:            false,                 
	}
}


func (g *GameState) CheckWin() {
	
	for _, char := range g.Word {
		found := false
		
		for _, guessed := range g.GuessedLetters {
			if guessed == char {
				found = true
				break  
			}
		}
		
		if !found {
			return
		}
	}
	
	
	g.Won = true      
	g.GameOver = true 
}




func (g *GameState) Guess(letter rune) (bool, error) {
	
	if g.GameOver {
		return false, fmt.Errorf("игра завершена")
	}


	letter = unicode.ToUpper(letter)
	
	
	for _, guessed := range g.GuessedLetters {
		if guessed == letter {
			return false, fmt.Errorf("буква '%c' уже угадана", letter)
		}
	}

	
	g.GuessedLetters = append(g.GuessedLetters, letter)

	
	found := false
	for _, char := range g.Word {
		if char == letter {
			found = true
			break  
		}
	}


	if !found {
		g.WrongAttempts++  
		
		
		if g.WrongAttempts >= g.MaxAttempts {
			g.GameOver = true  
		}
		return false, nil  
	}


	g.CheckWin()

	return true, nil  
}



func (g *GameState) GetCurrentState() string {
	result := ""

	for _, char := range g.Word {
		guessed := false
		
		for _, guessedChar := range g.GuessedLetters {
			if guessedChar == char {
				guessed = true
				break
			}
		}
	
		if guessed {
			result += string(char)
		} else {
			result += "*"
		}
	}
	return result
}


func (g *GameState) GetRemainingAttempts() int {
	return g.MaxAttempts - g.WrongAttempts
}


func (g *GameState) GetWord() string {
	return g.Word
}


func (g *GameState) IsGameOver() bool {
	return g.GameOver
}


func (g *GameState) IsWon() bool {
	return g.Won
}


func (g *GameState) GetGuessedLetters() []rune {
	return g.GuessedLetters
}


func (g *GameState) GetWrongAttempts() int {
	return g.WrongAttempts
}
