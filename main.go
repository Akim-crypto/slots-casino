package main

import "fmt"

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


func main(){
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