package channelsexample2

import "time"

func worker(done chan bool) {
	println("Working...")
	time.Sleep(time.Second)
	println("Done")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}