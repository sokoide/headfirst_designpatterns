package coffee

import (
	log "github.com/sirupsen/logrus"
)

// types
type HouseBlend struct {
	Beverage
}

type DarkRoast struct {
	Beverage
}

type Espresso struct {
	Beverage
}

type Decaf struct {
	Beverage
}

type Milk struct {
	CondimentDecorator
}

type Mocha struct {
	CondimentDecorator
}

type Soy struct {
	CondimentDecorator
}

type Whip struct {
	CondimentDecorator
}

// interfaces
type Beverage interface {
	Description() string
	Cost() float64
}

type CondimentDecorator struct {
	beverage Beverage
}

// mandate 'Beverag' implementation
var _ Beverage = (*Milk)(nil)
var _ Beverage = (*Mocha)(nil)
var _ Beverage = (*Soy)(nil)
var _ Beverage = (*Whip)(nil)

// methods
func (c *HouseBlend) Description() string { return "House Blend" }
func (c *DarkRoast) Description() string  { return "Dark Roast" }
func (c *Espresso) Description() string   { return "Espresso" }
func (c *Decaf) Description() string      { return "Decaf" }

func (c *HouseBlend) Cost() float64 { return 1.00 }
func (c *DarkRoast) Cost() float64  { return 2.00 }
func (c *Espresso) Cost() float64   { return 3.00 }
func (c *Decaf) Cost() float64      { return 4.00 }

func (c *Milk) Description() string  { return "Milk " + c.beverage.Description() }
func (c *Mocha) Description() string { return "Mocha" + c.beverage.Description() }
func (c *Soy) Description() string   { return "Soy" + c.beverage.Description() }
func (c *Whip) Description() string  { return "Whip" + c.beverage.Description() }

func (c *Milk) Cost() float64  { return 0.10 + c.beverage.Cost() }
func (c *Mocha) Cost() float64 { return 0.20 + c.beverage.Cost() }
func (c *Soy) Cost() float64   { return 0.30 + c.beverage.Cost() }
func (c *Whip) Cost() float64  { return 0.40 + c.beverage.Cost() }

func Run() {
	log.Infoln("--- Coffee ---")
	c := &HouseBlend{}
	m := &Milk{CondimentDecorator{c}}

	log.Infof("c: %s, %f\n", c.Description(), c.Cost())
	log.Infof("m: %s, %f\n", m.Description(), m.Cost())

	w := &Whip{CondimentDecorator{m}}
	log.Infof("w: %s, %f\n", w.Description(), w.Cost())

	mc := &Mocha{CondimentDecorator{w}}
	log.Infof("mc: %s, %f\n", mc.Description(), mc.Cost())

	s := &Soy{CondimentDecorator{mc}}
	log.Infof("s: %s, %f\n", "hoge", s.Cost())
}
