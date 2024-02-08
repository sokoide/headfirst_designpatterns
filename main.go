package main

import (
	coffee "github.com/sokoide/designpatterns/pkg/coffee"
	duck "github.com/sokoide/designpatterns/pkg/duck"
	pizzav2 "github.com/sokoide/designpatterns/pkg/pizza/v2"
	pizza "github.com/sokoide/designpatterns/pkg/pizza"
	weather "github.com/sokoide/designpatterns/pkg/weather"
)

func main() {
	duck.Run()
	weather.Run()
	coffee.Run()
	pizza.Run()
	pizzav2.Run()
}
