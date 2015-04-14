package main

func readOnlyChannel() <-chan string {
	channel := make(chan string)

	go func() {
		channel <- "Hello"
		channel <- "World"
		defer close(channel)
	}()

	return channel
}

func main() {
	hello := readOnlyChannel()

	println(<-hello)
	println(<-hello)
	println(<-hello)

	hello <- "Foobar"
}
