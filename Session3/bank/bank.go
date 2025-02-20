package bank

import "fmt"

type Bank struct {
	logs string
}

type BankAccount struct {
	accountHolder string
	balance       int
}

var accounts = make(map[*Bank]map[string]*BankAccount)


func NewBank() *Bank {
	bank := new(Bank)
	accounts[bank] = make(map[string]*BankAccount)
	return bank
}


func (bank *Bank) CreateBankAccount(accountHolder string, initialBalance int) error {
	if len(accountHolder) < 1 {
		return fmt.Errorf("Account holder name can not be empty")
	}
	if _, ok := accounts[bank][accountHolder]; ok{
		return fmt.Errorf("Account with this name Already exists")
	}

	if initialBalance < 0 {
		return fmt.Errorf("Account initial balance can not be negative")
	}

	accounts[bank][accountHolder] = &BankAccount{accountHolder: accountHolder, balance: initialBalance}

	return nil
}

func (bank *Bank) Deposit(accountHolder string, amount int) error {
	if _, ok := accounts[bank][accountHolder]; !ok {
		return fmt.Errorf("Account Does not Exists")
	}
	
	if amount < 1 {
		return fmt.Errorf("Amount Can not be zero or less")
	}

	accounts[bank][accountHolder].balance += amount
	bank.logs += fmt.Sprintf("Deposit %d from %s", amount, accountHolder + "\n")
	return nil
}

func (bank *Bank) Withdraw(accountHolder string, amount int) error {
	if _, ok := accounts[bank][accountHolder]; !ok {
		return fmt.Errorf("Account Does not Exists")
	}

	if amount < 1 {
		return fmt.Errorf("Amount Can not be zero or less")
	}

	if amount > accounts[bank][accountHolder].balance {
		return fmt.Errorf("Requested Amount can not be more than account balance")
	}

	accounts[bank][accountHolder].balance -= amount

	bank.logs += fmt.Sprintf("Withdraw %d from %s", amount, accountHolder + "\n")
	return nil
}

func (bank *Bank) GetBalance(accountHolder string) (int, error) {
	if _, ok := accounts[bank][accountHolder]; !ok {
		return 0, fmt.Errorf("Account Does not Exists")
	}
	return accounts[bank][accountHolder].balance, nil
}

func (bank *Bank) Transfer(from string, to string, amount int) error {
	if _, ok := accounts[bank][from]; !ok {
		return fmt.Errorf("Both Accounts must Exist")
	}
	if _, ok := accounts[bank][to]; !ok {
		return fmt.Errorf("Both Accounts must Exist")
	}

	if amount < 1 {
		return fmt.Errorf("Amount Can not be zero or less")
	}

	if amount > accounts[bank][from].balance {
		return fmt.Errorf("from account does not have enough balance")
	}
	accounts[bank][from].balance -= amount
	accounts[bank][to].balance += amount


	bank.logs += fmt.Sprintf("Transfer %d from %s to %s", amount, from, to + "\n")

	return nil
}

func (bank *Bank) TransactionLogs() string {
	if bank == nil {
		return ""
	}
	str := bank.logs
	if len(str) >= 1{
		return str[:len(str) - 1]
	} else {
		return ""
	}
}