package main

import (
	duck "github.com/sokoide/designpatterns/pkg/duck"
	weather "github.com/sokoide/designpatterns/pkg/weather"
)

func main() {
	duck.Run()
	weather.Run()
}
