package main

import (
	"fmt"
	"sync"
)

var (
	seats [39]bool
	mu    sync.Mutex
)

func bookSeat(user string, seatNumber int, wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	if seats[seatNumber] {
		fmt.Println(user, "- место", seatNumber, "УЖЕ ЗАНЯТО")
		return
	}

	seats[seatNumber] = true
	fmt.Println(user, "- УСПЕШНО забронировал место", seatNumber)
}

func main() {
	var wg sync.WaitGroup

	requests := []struct {
		user string
		seat int
	}{
		{"Анна", 5},
		{"Борис", 5},
		{"Виктор", 5},
		{"Дарья", 12},
		{"Егор", 12},
		{"Жанна", 38},
	}

	for _, req := range requests {
		wg.Add(1)
		go bookSeat(req.user, req.seat, &wg)
	}

	wg.Wait()
	fmt.Println("Все заявки обработаны.")
}
