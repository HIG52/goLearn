package accounts

import (
	"errors"
	"fmt"
)

// Account struct 클래스? 변수 모임?
type account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("돈이 부족합니다.")

// NewAccount create account
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

// (a account) 이걸 receiver라고 함 go에서는 이게method랑 동등한 관계에 있는거임
func (a *account) Deposit(amount int) {
	a.balance += amount
}

func (a account) Balance() int {
	return a.balance
}

func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil //error 에게도 두개의 value가 있는데 위에서 리턴해주는 에러와 nil이다 nil은 파이썬이나 자바스크립트의 null이나none 같은것이다
}

// change owner
func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a account) Owner() string {
	return a.owner
}

func (a account) String() string {
	return fmt.Sprint(a.owner, "'s account. \n Has: ", a.balance)
}
