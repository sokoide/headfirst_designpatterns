package pizza

import (
	"reflect"

	log "github.com/sirupsen/logrus"
)

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

func orderPizza(ty string) Pizza {
	var p Pizza
	if ty == "cheese" {
		p = &CheesePizza{BasePizza{"cheese"}}
	} else {
		p = &VeggiePizza{BasePizza{"veggie"}}
	}

	p.Prepare()
	p.Bake()
	p.Cut()
	p.Box()
	return p
}

func Run() {
	log.Infoln("-- Pizza --")
	p1 := orderPizza("cheese")
	p2 := orderPizza("veggie")
	log.Infof("p1: %+v", p1)
	log.Infof("p2: %+v", p2)

	// get type
	log.Infof("p1 type: %s:", reflect.TypeOf(p1))
	log.Infof("p2 type: %s:", reflect.TypeOf(p2))
}
