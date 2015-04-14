package main

func main() {
	channel := make(chan string)
	go func() {
		channel <- "Hello"
		channel <- "World"
		close(channel)
	}()

	// println(<-channel)
	// println(<-channel)
	// println(<-channel)

	for message := range channel {
		println("out: " + message)
	}
}
