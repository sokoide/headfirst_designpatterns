package main

import (
	coffee "github.com/sokoide/designpatterns/pkg/coffee"
	duck "github.com/sokoide/designpatterns/pkg/duck"
	pizza "github.com/sokoide/designpatterns/pkg/pizza"
	pizza_v2 "github.com/sokoide/designpatterns/pkg/pizza/v2"
	pizza_v3 "github.com/sokoide/designpatterns/pkg/pizza/v3"
	weather "github.com/sokoide/designpatterns/pkg/weather"
)

func main() {
	duck.Run()
	weather.Run()
	coffee.Run()
	pizza.Run()
	pizza_v2.Run()
	pizza_v3.Run()
}
