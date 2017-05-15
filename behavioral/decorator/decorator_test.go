package decorator

import "testing"

func TestEspressoWithMocha(t *testing.T) {
	var espressoWithMocha Beverage

	mocha := new(Mocha)
	mocha.beverage = new(Espresso)
	mocha.cost = 0.99
	mocha.description = "Mocha"

	espressoWithMocha = mocha

	beverageSum := (new(Espresso)).GetCost() + mocha.cost
	if espressoWithMocha.GetCost() != beverageSum {
		t.Errorf(
			"Wrong total espresso + mocha sum. Expected %s, but got %s",
			espressoWithMocha.GetCost(),
			beverageSum,
		)
	}

	beverageDesc := (new(Espresso)).GetDescription() + ", " + mocha.description
	if espressoWithMocha.GetDescription() != beverageDesc {
		t.Errorf(
			"Wrong description espresso + mocha. Expected %s, but got %s",
			espressoWithMocha.GetDescription(),
			beverageDesc,
		)
	}
}
