package strategy

const cantFlyMsg = "I can't fly!"
const flyingMsg = "I'm flying!"
const silenceMsg = "<<<Silence>>>"
const quackingMsg = "Quack!"

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
	return cantFlyMsg
}

type FlyWithWings struct{}

func (f *FlyWithWings) fly() string {
	return flyingMsg
}

type MuteQuack struct{}

func (m *MuteQuack) quack() string {
	return silenceMsg
}

type Quack struct{}

func (q *Quack) quack() string {
	return quackingMsg
}
