package goroutine

import (
	"sync"

	"github.com/bradhe/stopwatch"
	"github.com/gofiber/fiber/v2"
)

var wg sync.WaitGroup
var once sync.Once

type Counter struct {
	Value int
	mutex sync.Mutex
}

func (c *Counter) incrementValue() {
	c.mutex.Lock()
	c.Value++
	println("Count -", c.getValue())
	once.Do(func() {
		println("-=-=-=-=-=-=- [ ONCE ] -=-=-=-=-=-=-")
	})
	c.mutex.Unlock()
}

func (c *Counter) getValue() int {
	return c.Value
}

func loopIncrements(c *Counter, loop int) {
	defer wg.Done()

	for i := 0; i < loop; i++ {
		c.incrementValue()
	}
}

func TestMuTexLock(context *fiber.Ctx) error {

	w := stopwatch.Start()
	c := new(Counter)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go loopIncrements(c, 100)
	}

	wg.Wait()

	// loopIncrements(c, 1000000)

	result := c.getValue()
	println("Result =", result)

	w.Stop()

	return context.JSON(fiber.Map{
		"result": result,
		"timer":  w.Milliseconds(),
	})
}
