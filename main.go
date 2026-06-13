package main

import "fmt"

type Account struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		a.Balance += amount
		fmt.Println("Deposit Successful")
		fmt.Println("Amount Deposited:", amount)
		fmt.Println("New Balance:", a.Balance)
	} else {
		fmt.Println("Deposit amount must be positive")
	}
}

func (a *Account) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Withdrawal amount must be positive")
		return
	}

	if amount > a.Balance {
		fmt.Println("Insufficient Balance")
		return
	}

	a.Balance -= amount

	fmt.Println("Withdraw Successful")
	fmt.Println("Amount Withdrawn:", amount)
	fmt.Println("New Balance:", a.Balance)
}

func (a Account) CheckBalance() {
	fmt.Println("Current Balance:", a.Balance)
}

func (a Account) DisplayAccount() {
	fmt.Println("------ Account Information ------")
	fmt.Println("Account Number:", a.AccountNumber)
	fmt.Println("Holder Name:", a.HolderName)
	fmt.Println("Balance:", a.Balance)
	fmt.Println("---------------------------------")
}

func main() {
	account := Account{
		AccountNumber: "123456789",
		HolderName:    "John Doe",
		Balance:       1000,
	}

	account.DisplayAccount()

	account.CheckBalance()

	account.Deposit(500)

	account.Withdraw(200)

	account.Withdraw(1500)

	account.CheckBalance()
}




// package main

// import "fmt"

// type Account struct {
// 	AccountNumber string
// 	HolderName    string
// 	Balance       float64
// }

// func (a *Account) Deposit(amount float64) {
// 	if amount > 0 {
// 		a.Balance += amount
// 		fmt.Printf("Deposited %.2f to account %s. New balance: %.2f\n", amount, a.AccountNumber, a.Balance)
// 	} else {
// 		fmt.Println("Deposit amount must be positive.")
// 	}
// }

// func (a *Account) Withdraw(amount float64) {
// 	if amount > 0 {
// 		if a.Balance >= amount {
// 			a.Balance -= amount
// 			fmt.Printf("Withdrew %.2f from account %s. New balance: %.2f\n", amount, a.AccountNumber, a.Balance)
// 		} else {
// 			fmt.Println("Insufficient funds.")
// 		}
// 	} else {
// 		fmt.Println("Withdrawal amount must be positive.")
// 	}
// }

// func (a *Account) MainBlance() {
// 	fmt.Printf("Current balance for account %s: %.2f\n", a.AccountNumber, a.Balance)
// }

// func main() {
// 	account := Account{
// 		AccountNumber: "123456789",
// 		HolderName:    "John Doe",
// 		Balance:       1000.00,
// 	}

// 	account.MainBlance()
// 	account.Deposit(500.00)
// 	account.Withdraw(200.00)
// 	account.Withdraw(1500.00)
// 	account.MainBlance()
// }