package main

import "testing"

func TestDuck(t *testing.T) {
	d := &Duck{&FlyNoWay{}, &MuteQuack{}}

	if d.fly() != "I can't fly!" {
		t.Error("Fly error")
	}

	if d.quack() != "<<<Silence>>>" {
		t.Error("Quack error")
	}

	d.SetFlyable(&FlyWithWings{})
	d.SetQuackable(&Quack{})

	if d.fly() != "I'm flying!" {
		t.Error("Fly error")
	}

	if d.quack() != "Quack!" {
		t.Error("Quack error")
	}
}
