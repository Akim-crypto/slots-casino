package main

import "fmt"

func getName() string {
    name := ""

    fmt.Println("Welcome to Akim`s Casino")...
    fmt.Printlf("Enter your name:")
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
    }
}


func main(){
    balance := 200

    getName()
}