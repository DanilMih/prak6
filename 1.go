package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mut sync.Mutex
	views := 0

	gorutinesCount := 50
	viewFromUser := 1000

	for i := 0; i < gorutinesCount; i++ {
		wg.Add(1)
		go worker(&views, &mut, &wg, viewFromUser)
	}

	wg.Wait()
	fmt.Println("Получено:", views, "просмотров")

}

func worker(views *int, mut *sync.Mutex, wg *sync.WaitGroup, iter int) {

	defer wg.Done()

	for i := 0; i < iter; i++ {
		mut.Lock()
		*views++
		mut.Unlock()
	}
}
