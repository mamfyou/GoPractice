package bank

import "testing"

func TestCreateInitialBalance_NegativeAmount(t *testing.T) {
	b := NewBank()

	err := b.CreateBankAccount("ali", -100)
	if err == nil {
		t.Errorf("expected err but got nil")
	}
}

func TestCreateDuplicateBankAccount(t *testing.T) {
	b := NewBank()

	err := b.CreateBankAccount("ali", 100)
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	err = b.CreateBankAccount("ali", 1)
	if err == nil {
		t.Errorf("expected error but got nil")
	}
}

func TestWidraw(t *testing.T) {
	b := NewBank()

	err := b.CreateBankAccount("ali", 100)
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	err = b.Withdraw("ali", 50)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	balance, err := b.GetBalance("ali")
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	if balance != 50 {
		t.Errorf("expected 50, got %d", balance)
	}
}

func TestWithdraw_NegativeAmount(t *testing.T) {
	b := NewBank()

	err := b.CreateBankAccount("ali", 100)
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	err = b.Withdraw("ali", -50)
	if err == nil {
		t.Errorf("expected error but got nil")
	}
}

func TestWithdraw_ZeroAmount(t *testing.T) {
	b := NewBank()

	err := b.CreateBankAccount("ali", 100)
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	err = b.Withdraw("ali", 0)
	if err == nil {
		t.Errorf("expected error but got nil")
	}
}

func TestWithdraw_FullBalance(t *testing.T) {
	b := NewBank()

	err := b.CreateBankAccount("ali", 100)
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	err = b.Withdraw("ali", 100)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	balance, err := b.GetBalance("ali")
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	if balance != 0 {
		t.Errorf("expected 0, got %d", balance)
	}
}

func TestWithdraw_MissingAccount(t *testing.T) {
	b := NewBank()

	err := b.Withdraw("non_existent", 50)
	if err == nil {
		t.Errorf("expected error for missing account, got nil")
	}
}

func TestWithdraw_InsufficientFunds(t *testing.T) {
	b := NewBank()

	err := b.CreateBankAccount("ali", 100)
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	err = b.Withdraw("ali", 150)
	if err == nil {
		t.Errorf("expected to get error but got nil")
	}

	balance, err := b.GetBalance("ali")
	if err != nil {
		t.Errorf("execpted error: %v", err)
	}

	if balance != 100 {
		t.Errorf("expected 100 but got %d", balance)
	}
}

func TestDeposit(t *testing.T) {
	b := NewBank()

	err := b.CreateBankAccount("ali", 100)
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	err = b.Deposit("ali", 50)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	balance, err := b.GetBalance("ali")
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	if balance != 150 {
		t.Errorf("expected 150.00, got %d", balance)
	}
}

func TestDeposit_MissingAccount(t *testing.T) {
	b := NewBank()

	err := b.Deposit("non_existent", 50)
	if err == nil {
		t.Errorf("expected error for missing account, got nil")
	}
}

func TestDeposit_NegativeAmount(t *testing.T) {
	b := NewBank()

	err := b.CreateBankAccount("ali", 100)
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	err = b.Deposit("ali", -50)
	if err == nil {
		t.Errorf("expected error but got nil")
	}
}

func TestDeposit_ZeroAmount(t *testing.T) {
	b := NewBank()

	err := b.CreateBankAccount("ali", 100)
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	err = b.Deposit("ali", 0)
	if err == nil {
		t.Errorf("expected error but got nil")
	}
}

func TestTransfer(t *testing.T) {
	b := NewBank()

	err := b.CreateBankAccount("ali", 100)
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	err = b.CreateBankAccount("fatemeh", 200)
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	err = b.Transfer("ali", "fatemeh", 25)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	aliBalance, err := b.GetBalance("ali")
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	if aliBalance != 75 {
		t.Errorf("expected 75, got %d", aliBalance)
	}

	fatemehBalance, err := b.GetBalance("fatemeh")
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	if fatemehBalance != 225 {
		t.Errorf("expected 225, got %d", fatemehBalance)
	}
}

func TestTransactionLogs(t *testing.T) {
	b := NewBank()

	b.CreateBankAccount("ali", 100)
	b.CreateBankAccount("fatemeh", 45)

	err := b.Withdraw("ali", 50)
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	err = b.Deposit("fatemeh", 4)
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	b.Transfer("ali", "fatemeh", 8)

	expected := "Withdraw 50 from ali\nDeposit 4 from fatemeh\nTransfer 8 from ali to fatemeh"

	if b.TransactionLogs() != expected {
		t.Errorf("expected %s but got %s", expected, b.TransactionLogs())
	}

}

func TestTransactionLogs_NilBankInstance(t *testing.T) {
	var b *Bank = nil

	if b.TransactionLogs() != "" {
		t.Errorf("expected empty string but got %s", b.TransactionLogs())
	}
}

func TestTransfer_MissingFromAccount(t *testing.T) {
	b := NewBank()

	err := b.CreateBankAccount("to", 100)
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	err = b.Transfer("missing_from", "to", 50)
	if err == nil {
		t.Errorf("Expected error for missing 'from' account but got nil")
	}
}

func TestTransfer_MissingToAccount(t *testing.T) {
	b := NewBank()

	err := b.CreateBankAccount("from", 100)
	if err != nil {
		t.Errorf("unexpted error: %v", err)
	}

	err = b.Transfer("from", "missing_to", 50)
	if err == nil {
		t.Errorf("Expected error for missing 'to' account but got nil")
	}
}

func TestCreateBankAccount_EmptyAccountHolder(t *testing.T) {
	b := NewBank()

	err := b.CreateBankAccount("", 100)
	if err == nil {
		t.Errorf("expected err but got nil")
	}
}
