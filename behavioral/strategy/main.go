package main

type Duck struct {
	Flyable
	Quackable
}

type Flyable interface {
	fly() string
}

type Quackable interface {
	quack() string
}

func (d *Duck) PerformFly() string {
	return d.fly()
}

func (d *Duck) PerformQuack() string {
	return d.quack()
}

func (d *Duck) SetFlyable(flyBehavior Flyable) {
	d.Flyable = flyBehavior
}

func (d *Duck) SetQuackable(quackBehavior Quackable) {
	d.Quackable = quackBehavior
}

type FlyNoWay struct{}

func (f *FlyNoWay) fly() string {
	return "I can't fly!"
}

type FlyWithWings struct{}

func (f *FlyWithWings) fly() string {
	return "I'm flying!"
}

type MuteQuack struct{}

func (m *MuteQuack) quack() string {
	return "<<<Silence>>>"
}

type Quack struct{}

func (q *Quack) quack() string {
	return "Quack!"
}
