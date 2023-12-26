package duck

import (
	log "github.com/sirupsen/logrus"
)

// types
type Duck struct {
}

type MallarDuck struct {
	Duck
	FlyWithWings
	RegularQuack
}

type RedheadDuck struct {
	Duck
	FlyWithWings
	RegularQuack
}

type RubberDuck struct {
	Duck
	FlyNoWay
	Squeak
}

type FlyWithWings struct{}
type FlyNoWay struct{}
type RegularQuack struct{}
type Squeak struct{}
type MuteQuack struct{}

// interfaces
type FlyBehavior interface {
	Fly()
}

type QuackBehavior interface {
	Quack()
}

func (f *FlyWithWings) Fly() {
	log.Infoln("* FlyWithWings")
}

func (f *FlyNoWay) Fly() {
	log.Infoln("* FlyNoWay")
}

func (q *RegularQuack) Quack() {
	log.Infoln("* Quack")
}

func (q *Squeak) Quack() {
	log.Infoln("* Squeak")
}

func (q *MuteQuack) Quack() {
}

// methods
func (d *Duck) Display() {
	log.Infoln("Duck")
}

func (d *MallarDuck) Display() {
	log.Infoln("MallarDuck")
}

func (d *RedheadDuck) Display() {
	log.Infoln("RedheadDuck")
}

func (d *RubberDuck) Display() {
	log.Infoln("RubberDuck")
}

// functions
func Run() {
	log.Infoln("--- Duck ---")
	red := &RedheadDuck{}
	red.Display()
	red.Fly()
	red.Quack()

	rubber := &RubberDuck{}
	rubber.Display()
	rubber.Fly()
	rubber.Quack()
}
