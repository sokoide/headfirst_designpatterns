package pizza

import (
	"reflect"

	log "github.com/sirupsen/logrus"
)

// types
type PizzaStore struct {
	factory *SimpleFactory
}

type SimpleFactory struct{}

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

// methods
func (p *BasePizza) Prepare() {
}

func (p *BasePizza) Bake() {
}

func (p *BasePizza) Cut() {
}

func (p *BasePizza) Box() {
}

func (s *PizzaStore) orderPizza(ty string) Pizza {
	p := s.factory.CreatePizza(ty)

	p.Prepare()
	p.Bake()
	p.Cut()
	p.Box()
	return p
}

func (f *SimpleFactory) CreatePizza(ty string) Pizza {
	var p Pizza
	if ty == "cheese" {
		p = &CheesePizza{BasePizza{"cheese"}}
	} else {
		p = &VeggiePizza{BasePizza{"veggie"}}
	}

	return p
}

// funtions
func Run() {
	log.Infoln("-- Pizza --")
	store := PizzaStore{factory: &SimpleFactory{}}

	p1 := store.orderPizza("cheese")
	log.Infof("p1: %+v", p1)

	p2 := store.orderPizza("veggie")
	log.Infof("p2: %+v", p2)

	// type assertion
	if value, ok := p2.(*CheesePizza); ok {
		log.Infof("if1 value: %+v", value)
	} else {
		log.Infoln("else1")
	}
	if value, ok := p2.(*VeggiePizza); ok {
		log.Infof("if2 value: %+v", value)
	} else {
		log.Infoln("else2")
	}
	// get type
	log.Infoln(reflect.TypeOf(p2))
}
