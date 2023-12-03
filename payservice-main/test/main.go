package main

import (
	"fmt"
	"sync"
)

type PaymentService struct {
	balance map[string]float64
	mu      sync.Mutex
}

func NewPaymentService() *PaymentService {
	return &PaymentService{
		balance: make(map[string]float64),
	}
}

func (p *PaymentService) Deposit(accountID string, amount float64) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.balance[accountID] += amount
	fmt.Printf("Deposited %.2f into account %s\n", amount, accountID)
}

func (p *PaymentService) Withdraw(accountID string, amount float64) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if balance, ok := p.balance[accountID]; ok && balance >= amount {
		p.balance[accountID] -= amount
		fmt.Printf("Withdrawn %.2f from account %s\n", amount, accountID)
	} else {
		fmt.Printf("Insufficient funds for withdrawal from account %s\n", accountID)
	}
}

func (p *PaymentService) Transfer(senderID, receiverID string, amount float64, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	p.mu.Lock()
	defer p.mu.Unlock()

	if balance, ok := p.balance[senderID]; ok && balance >= amount {
		p.balance[senderID] -= amount
		p.balance[receiverID] += amount
		ch <- fmt.Sprintf("Transferred %.2f from %s to %s", amount, senderID, receiverID)
	} else {
		ch <- fmt.Sprintf("Insufficient funds for transfer from %s to %s", senderID, receiverID)
	}
}

func main() {
	paymentService := NewPaymentService()

	paymentService.Deposit("account1", 100.0)
	paymentService.Deposit("account2", 50.0)

	var wg sync.WaitGroup
	ch := make(chan string, 2)

	// Горутина для перевода средств

	wg.Add(4)
	go paymentService.Transfer("account2", "account1", 20.0, &wg, ch)
	go paymentService.Transfer("account2", "account1", 20.0, &wg, ch)
	go paymentService.Transfer("account2", "account1", 20.0, &wg, ch)
	go paymentService.Transfer("account2", "account1", 20.0, &wg, ch)

	// Закрытие канала после завершения всех горутин
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Чтение из канала
	for message := range ch {
		fmt.Println(message)
	}

	// Вывод остатка на счетах
	fmt.Printf("Balance of account1: %.2f\n", paymentService.balance["account1"])
	fmt.Printf("Balance of account2: %.2f\n", paymentService.balance["account2"])
}
