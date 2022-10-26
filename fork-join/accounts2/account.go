package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Account struct {
	balance int64
}
type Transaction struct {
	From   *Account
	To     *Account
	amount int64
}

type SharedContext struct {
	transactions []*Transaction
	group        *sync.WaitGroup
	flag         *int64
}

func transfer(flag *int64, transaction *Transaction) {

	from := transaction.From
	to := transaction.To
	amount := transaction.amount

	// Question: Is this okay to do? with having the atomics happen inside the the if-statement
	if from.balance >= amount {
		atomic.AddInt64(&(from.balance), -amount)
		atomic.AddInt64(&(to.balance), amount)
	}

}

func worker(start int, context *SharedContext) {

	for i := start; i < start+40; i++ {
		transaction := context.transactions[i]
		transfer(context.flag, transaction)
	}
	context.group.Done()

}
func forkJoin(context *SharedContext) {

	/*
	  For this example of Fork-Join we will hardcode in the threshold
	  and how many threads to correctly spawn based on the number of
	  transactions being equal to 200. However, you will need to think
	  how you can do this dynamically.
	*/

	//Sequential condition to perform the computation sequentially
	//This is very specific trivial case here for only this specific example.
	//In general you should basis this off some sequential threshold value.
	if len(context.transactions) != 200 {
		fmt.Printf("Got here:=%v", len(context.transactions))
	} else {
		for i := 0; i < 160; i += 40 {
			context.group.Add(1)
			go worker(i, context)
		}
		context.group.Add(1)
		worker(160, context)
		context.group.Wait()
	}
}

func main() {

	//Create accounts
	account1 := &Account{100}
	account2 := &Account{100}

	//Initialize all transactions
	transactions := make([]*Transaction, 0)

	for i := 0; i < 100; i++ {
		transactions = append(transactions, &Transaction{account1, account2, 1})
		transactions = append(transactions, &Transaction{account2, account1, 1})
	}

	group := &sync.WaitGroup{}
	var flag int64
	context := &SharedContext{transactions, group, &flag}
	forkJoin(context)

	fmt.Println("Account1 Balance = $", account1.balance)
	fmt.Println("Account2 Balance = $", account2.balance)

}
