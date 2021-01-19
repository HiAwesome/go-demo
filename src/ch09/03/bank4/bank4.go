// 使用多读单写锁执行转账事务
// @author moqi
// On 2021/1/19 22:06:43
package bank4

import "sync"

var (
	mu      sync.RWMutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance() int {
	mu.RLock()
	b := balance
	mu.RUnlock()
	return b
}
