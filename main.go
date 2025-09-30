package main

import (
	"math/rand"
	"fmt"
)

func getName() string {
    name := ""

    fmt.Println("Welcome to Akim`s Casino...")
    fmt.Printf("Enter your name:")
    _, err := fmt.Scanln(&name)
    if err != nil {
        return ""
    }
    fmt.Printf("Welcome %s , lets play\n",name)
    return name
}

func getBet(balance uint) uint {
    var bet uint
    for true {
        fmt.Printf("Enter your bet , or 0 to quit (balance = $%d): ",balance)
        fmt.Scan(&bet)

        if bet > balance {
            fmt.Println("Bet cannot be larger than bet")
        } else {
            break
        }
    }
    return bet
}

func generateSymbolArray(symbols map[string]uint) []string{
    symbolArr := []string{}
    for symbol , count := range symbols{
        for i := 0; i < int(count); i ++{
            symbolArr = append(symbolArr, symbol)
        }
    }
    return symbolArr
}


func getRandomNumber(min int,max int) int {
    randomNumber:= rand.Intn(max - min +1) + min
    return randomNumber
}


func getSpin(reel []string,rows int,cols int) [][]string {
    result := [][]string{}

    for i := 0; i < rows; i++ {
        result = append(result, []string{})
    }

    for col := 0 ; col < cols; col ++ {
        selected := map[int]bool{}
        for row := 0 ; row<rows ; row ++{
            for true {
                randomIndex := getRandomNumber(0, len(reel)-1)
                _,exists := selected[randomIndex]
                if !exists{
                    selected[randomIndex] = true
                    result[row] = append(result[row], reel[randomIndex])
                    break
                }
            }
        }
    }
    return result
}




func main(){
    symbols := map[string]uint{
        "A":4,
        "B":7,
        "C":12,
        "D":20,
    }
    // multipliers := map[string]uint{
    //     "A":20,
    //     "B":10,
    //     "C":5,
    //     "D":2,
    // }

    symbolArr := generateSymbolArray(symbols)
    spin := getSpin(symbolArr, 3,3)
    fmt.Println(spin)




    var balance uint = 200
    getName()

    for balance > 0 {
        bet := getBet(balance)
        if bet == 0 {
            fmt.Println("Your balance == 0 ")
            break
        }
        balance -= bet
    }
    
    fmt.Printf("You left with , $%d\n",balance)
   

}