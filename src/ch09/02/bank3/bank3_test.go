//
// @author moqi
// On 2021/1/19 00:09:31
package bank3_test

import (
	"moqi.com/go/ch09/02/bank3"
	"sync"
	"testing"
)

func TestBank(t *testing.T) {
	var n sync.WaitGroup

	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			bank3.Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := bank3.Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
