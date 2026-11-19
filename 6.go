package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func saveLog(workerID int, message string) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Printf("[Воркер %d] %s \n", workerID, message)
}

func runWorker(id int, stepsCount int, wg *sync.WaitGroup) {
	defer wg.Done()

	for step := 1; step <= stepsCount; step++ {
		msg := fmt.Sprintf("выполнил шаг %d из %d", step, stepsCount)
		saveLog(id, msg)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	const workerCount = 5
	const stepCount = 4

	var wg sync.WaitGroup

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go runWorker(i, stepCount, &wg)
	}

	wg.Wait()
	fmt.Println("Все воркеры завершили работу.")
}
