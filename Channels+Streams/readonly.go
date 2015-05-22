package main

func readOnlyChannel() <-chan string {
	channel := make(chan string)

	go func() {
		defer close(channel)
		channel <- "Hello"
		channel <- "World"
	}()

	return channel
}

func main() {
	hello := readOnlyChannel()

	println(<-hello)
	println(<-hello)

	// hello <- "Foobar"
}
