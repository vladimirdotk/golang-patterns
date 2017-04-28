package strategy

import "testing"

const errorFmt = "%s error! Expected %s, got %s"

func TestDuck(t *testing.T) {
	d := &Duck{&FlyNoWay{}, &MuteQuack{}}

	if fly := d.fly(); fly != cantFlyMsg {
		t.Errorf(errorFmt, "Fly", cantFlyMsg, fly)
	}

	if quack := d.quack(); quack != silenceMsg {
		t.Error(errorFmt, "Quack", silenceMsg, quack)
	}

	d.SetFlyable(&FlyWithWings{})
	d.SetQuackable(&Quack{})

	if fly := d.fly(); fly != flyingMsg {
		t.Error(errorFmt, "Fly", flyingMsg, fly)
	}

	if quack := d.quack(); quack != quackingMsg {
		t.Error(errorFmt, "Quack", quackingMsg, quack)
	}
}
