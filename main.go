package main

import (
    "fmt"
    "task-bank/bank"
)

func main() {
    account := bank.Account{}
    
    if err := account.Deposit(100.50); err != nil {
        fmt.Println("Error:", err)
    }
    
    if err := account.Withdraw(50); err != nil {
        fmt.Println("Error:", err)
    }
    
    if err := account.Withdraw(100); err != nil {
        fmt.Println("Error:", err)
    }

    fmt.Println("Final Balance:", account.GetBalance())
}