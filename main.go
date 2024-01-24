package main

import (
	coffee "github.com/sokoide/designpatterns/pkg/coffee"
	duck "github.com/sokoide/designpatterns/pkg/duck"
	pizza "github.com/sokoide/designpatterns/pkg/pizza"
	weather "github.com/sokoide/designpatterns/pkg/weather"
)

func main() {
	duck.Run()
	weather.Run()
	coffee.Run()
	pizza.Run()
}
