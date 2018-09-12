# Tick every x milliseconds, display message until it has been done y amount of times

## Code put together using this resource:
* https://play.golang.org/p/Ctg3_AQisl

You will need to get this package to make it work:
* `go get github.com/karolispx/golang-ticker`

## Sample usage:
* `ticker.Tick(100, "Hello world", 10)`

## Example:
```
package main

import "github.com/karolispx/golang-ticker"

func main() {
	ticker.Tick(100, "Hello world", 20)
}

```
