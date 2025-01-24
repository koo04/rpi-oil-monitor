package main

import "github.com/warthog618/go-gpiocdev"

func main() {
	l, _ := gpiocdev.RequestLine("gpiochip0", 4)
	_ = l
}
