package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	candidates := []string{"Кандидат А", "Кандидат Б", "Кандидат В"}
	results := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	const voters = 50
	const votesFromVoter = 20

	wg.Add(voters)

	for i := 0; i < voters; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < votesFromVoter; j++ {

				idx := rand.Intn(len(candidates))
				c := candidates[idx]

				mu.Lock()
				results[c]++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	total := 0
	for _, c := range candidates {
		fmt.Println(c+":", results[c])
		total += results[c]
	}

	fmt.Println("Всего:", total)
}
