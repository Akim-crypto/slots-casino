package main

import (
	"fmt"
	"math/rand"
)

func getName() string {
	name := ""
	fmt.Println("Welcome to Akim`s Casino...")
	fmt.Printf("Enter your name:")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s , lets play\n", name)
	return name
}

func getBet(balance uint) uint {
	var bet uint
	for {
		fmt.Printf("Enter your bet , or 0 to quit (balance = $%d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bet cannot be larger than balance")
		} else {
			break
		}
	}
	return bet
}

func generateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}
	for symbol, count := range symbols {
		for i := 0; i < int(count); i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}
	return symbolArr
}

func getRandomNumber(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func getSpin(reel []string, rows int, cols int) [][]string {
	result := [][]string{}
	for i := 0; i < rows; i++ {
		result = append(result, []string{})
	}
	for col := 0; col < cols; col++ {
		selected := map[int]bool{}
		for row := 0; row < rows; row++ {
			for {
				randomIndex := getRandomNumber(0, len(reel)-1)
				if !selected[randomIndex] {
					selected[randomIndex] = true
					result[row] = append(result[row], reel[randomIndex])
					break
				}
			}
		}
	}
	return result
}

func printSpin(spin [][]string) {
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Print(symbol)
			if j != len(row)-1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}

// <<< ВНЕ main >>>
func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}
	for _, row := range spin {
		win := true
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if checkSymbol != symbol {
				win = false
				break
			}
		}
		if win {
			lines = append(lines, multipliers[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return lines
}

func main() {
	symbols := map[string]uint{"A": 4, "B": 7, "C": 12, "D": 20}
	multipliers := map[string]uint{"A": 20, "B": 10, "C": 5, "D": 2}

	symbolArr := generateSymbolArray(symbols)

	var balance uint = 200
	getName()

	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			fmt.Println("You quit the game.")
			break
		}
		balance -= bet
		spin := getSpin(symbolArr, 3, 3)
		printSpin(spin)
		winningLines := checkWin(spin, multipliers)
		for i , multi := range winningLines{
            win := multi * bet
            balance += win
            if multi > 0 {
                fmt.Printf("Won $%d, (%dx) on line #%d\n",win,multi , i +1)
            }
        }
	}

	fmt.Printf("You left with , $%d\n", balance)
}
