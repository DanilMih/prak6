package main

import (
	"fmt"
	"sync"
)

var (
	stock = map[string]int{
		"Ноутбук": 5,
		"Мышь":    10,
	}
	mu sync.Mutex
)

func addProduct(name string, count int, wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	stock[name] += count
	fmt.Println("Поступление:", name, "+", count, " Остаток:", stock[name])
}

func sellProduct(name string, count int, wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	if stock[name] < count {
		fmt.Println("Отказ в продаже:", name, count, "шт. (недостаточно на складе, есть только", stock[name], ")")
		return
	}

	stock[name] -= count
	fmt.Println("Продажа:", name, "-", count, " Остаток:", stock[name])
}

func main() {
	var wg sync.WaitGroup

	wg.Add(6)
	go sellProduct("Ноутбук", 3, &wg)
	go sellProduct("Ноутбук", 4, &wg)
	go addProduct("Ноутбук", 5, &wg)
	go sellProduct("Мышь", 6, &wg)
	go addProduct("Мышь", 10, &wg)
	go sellProduct("Мышь", 8, &wg)

	wg.Wait()
	fmt.Println("Итоговые остатки на складе:", stock)
}
