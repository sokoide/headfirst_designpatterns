package pizza

import (
	log "github.com/sirupsen/logrus"
)

// types
type BasePizzaStore struct {
	PizzaStore
}

type NYPizzaStore struct {
	BasePizzaStore
}

type ChicagoPizzaStore struct {
	BasePizzaStore
}

type BasePizza struct{ Name string }
type CheesePizza struct{ BasePizza }
type PepperoniPizza struct{ BasePizza }
type ClamPizza struct{ BasePizza }
type VeggiePizza struct{ BasePizza }

// interfaces
type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type PizzaStore interface {
	CreatePizza(ty string) Pizza
}

// methods
func (p *BasePizza) Prepare() {
}

func (p *BasePizza) Bake() {
}

func (p *BasePizza) Cut() {
}

func (p *BasePizza) Box() {
}

func (s *BasePizzaStore) orderPizza(ty string) Pizza {
	p := s.CreatePizza(ty)

	p.Prepare()
	p.Bake()
	p.Cut()
	p.Box()
	return p
}

func (s *NYPizzaStore) CreatePizza(ty string) Pizza {
	var p Pizza
	if ty == "cheese" {
		p = &CheesePizza{BasePizza{"NY cheese"}}
	} else {
		p = &VeggiePizza{BasePizza{"NY veggie"}}
	}

	return p
}

func (s *ChicagoPizzaStore) CreatePizza(ty string) Pizza {
	var p Pizza
	if ty == "cheese" {
		p = &CheesePizza{BasePizza{"Chicago cheese"}}
	} else {
		p = &VeggiePizza{BasePizza{"Chicago veggie"}}
	}

	return p
}

// funtions
func Run() {
	log.Infoln("-- Pizza v3 --")
	var store BasePizzaStore
	store.PizzaStore = &NYPizzaStore{}

	p1 := store.orderPizza("cheese")
	log.Infof("p1: %+v", p1)

	store.PizzaStore = &ChicagoPizzaStore{}
	p2 := store.orderPizza("veggie")
	log.Infof("p2: %+v", p2)
}
