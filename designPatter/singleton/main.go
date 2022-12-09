package main

import "sync"

type singleton struct{}

var (
	instance     *singleton
	instanceOnce sync.Once
)

func main() {

}

func GetSingleton() *singleton {
	instanceOnce.Do(func() {
		instance = &singleton{}
	})
	return instance
}
