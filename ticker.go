package ticker

import (
	"fmt"
	"time"
)

/*
// Tick every x milliseconds, display message until it has been done y amount of times
// Usage: Tick(milliseconds, "Your message", times) [if times is 0 - it will tick until the program is stopped]
// Example: Tick(100, "Hello world", 10)
*/
func Tick(milliseconds int, message string, times int) {
	d := time.Duration(milliseconds) * time.Millisecond

	counter := 0

	for range time.Tick(d) {
		fmt.Println(message)

		counter++

		if times != 0 {
			if counter >= times {
				break
			}
		}
	}
}
