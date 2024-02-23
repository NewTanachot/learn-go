package goroutine

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Data string `json:"data"`
}

type PubSub struct {
	Subs  []chan Message
	Mutex sync.Mutex
}

func TestPubSub(ctx *fiber.Ctx) error {
	message := ctx.Params("message")
	pubsub := new(PubSub)

	pubsub.publish(&message)

	sub1 := pubsub.subscribe()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range *sub1 {
			println("Receive Message:", msg.Data)
		}
	}()
	wg.Wait()
	// sub2 := pubsub.subscribe()

	// go func() {
	// 	for msg := range *sub2 {
	// 		println("Receive Message:", msg.Data)
	// 	}
	// }()

	return ctx.JSON("Add message success")
}

func (ps *PubSub) subscribe() *chan Message {
	ps.Mutex.Lock()
	defer ps.Mutex.Unlock()

	channel := make(chan Message, 1)
	ps.Subs = append(ps.Subs, channel)

	return &channel
}

func (ps *PubSub) publish(text *string) {
	ps.Mutex.Lock()
	defer ps.Mutex.Unlock()

	for _, sub := range ps.Subs {
		sub <- Message{
			Data: *text,
		}
	}
}
