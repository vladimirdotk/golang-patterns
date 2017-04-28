package strategy

import "testing"

const errorFmt = "%s error! Expected %s, got %s"

func TestDuck(t *testing.T) {
	d := &Duck{&FlyNoWay{}, &MuteQuack{}}

	if fly := d.fly(); fly != cantFlyMsg {
		t.Errorf(errorFmt, "Fly", cantFlyMsg, fly)
	}

	if quack := d.quack(); quack != silenceMsg {
		t.Errorf(errorFmt, "Quack", silenceMsg, quack)
	}

	d.SetFlyable(&FlyWithWings{})
	d.SetQuackable(&Quack{})

	if fly := d.fly(); fly != flyingMsg {
		t.Errorf(errorFmt, "Fly", flyingMsg, fly)
	}

	if quack := d.quack(); quack != quackingMsg {
		t.Errorf(errorFmt, "Quack", quackingMsg, quack)
	}
}
