package goroutine

import (
	"github.com/gofiber/fiber/v2"
)

func TestChannel(context *fiber.Ctx) error {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go assignValueToChannel(channel1)
	go assignValueToChannel(channel2)

	isClose1, isClose2 := false, false

	for {

		if isClose1 && isClose2 {
			break
		}

		select {
		case v, ok := <-channel1:
			if !ok {
				isClose1 = !ok
				continue
			}
			println("channel1", v)
		case v, ok := <-channel2:
			if !ok {
				isClose2 = !ok
				continue
			}
			println("channel2", v)
		}
	}

	// for v := range channel1 {
	// 	println(v)
	// }

	// select {
	// case v := <-channel1:
	// 	println("channel1", v)
	// case v := <-channel2:
	// 	println("channel2", v)
	// default:
	// 	println("No message received")
	// }

	return context.SendStatus(fiber.StatusOK)
}

func assignValueToChannel(channel chan int) {
	channel <- 10
	close(channel)
}
