package goroutine

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

var mutex sync.Mutex
var waitgroup sync.WaitGroup

// BUG
func TestCond(context *fiber.Ctx) error {
	cond := sync.NewCond(&mutex)
	// channel := make(chan bool, 1)

	// channel <- false

	// waitgroup.Add(1)
	go waitTask(cond)
	// waitgroup.Wait()

	time.Sleep(time.Second)

	waitgroup.Add(1)
	mutex.Lock()

	// channel <- true
	// close(channel)

	cond.Signal()
	println("Send Signal!!!")

	mutex.Unlock()
	waitgroup.Wait()

	println("Work done.")

	return context.SendStatus(fiber.StatusOK)
}

func waitTask(cond *sync.Cond) {
	defer waitgroup.Done()

	mutex.Lock()

	println("Goroutine: waiting for condition...")
	// waitgroup.Done()

	// for value := range *channel {
	// 	println("channel value :", value)
	// 	for !value {
	// 		cond.Wait()
	// 	}
	// }

	cond.Wait()

	println("Goroutine: condition met. Processing...")
	mutex.Unlock()
}

// correct example
// func main() {
// 	var mutex sync.Mutex
// 	cond := sync.NewCond(&mutex)

// 	ready := make(chan bool) // Channel for signaling readiness
// 	wg := sync.WaitGroup{}

// 	// Launching goroutine
// 	wg.Add(1)
// 	go waitForCondition(&mutex, cond, ready, &wg)

// 	// Simulate some work (e.g., loading resources)
// 	// Here, you might perform actual work instead of sleeping
// 	// For demonstration, I'll use a sleep here
// 	time.Sleep(2 * time.Second)

// 	// Signal the condition
// 	mutex.Lock()
// 	ready <- true
// 	mutex.Unlock()
// 	fmt.Println("Signal sent!")

// 	wg.Wait() // Wait for the goroutine to complete
// 	fmt.Println("Main: Work is done.")
// }

// func waitForCondition(mutex *sync.Mutex, cond *sync.Cond, ready chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	fmt.Println("Goroutine: Waiting for the condition...")

// 	mutex.Lock()
// 	for !<-ready {
// 		cond.Wait() // Wait for the condition
// 	}
// 	fmt.Println("Goroutine: Condition met, proceeding...")
// 	mutex.Unlock()
// }
