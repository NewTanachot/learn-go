package goroutine

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

func TestWaitGroup(context *fiber.Ctx) error {
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each
	j := 5
	wg.Add(j)

	for i := 1; i <= j; i++ {
		go worker(i, &wg)
	}

	wg.Wait() // Block until the WaitGroup counter goes back to 0; all workers are done

	println("All workers completed")
	return context.SendStatus(fiber.StatusOK)
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	fmt.Printf("Worker %d starting\n", id)

	// Simulate some work by sleeping
	sleepDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(sleepDuration)

	fmt.Printf("Worker %d done\n", id)
}
