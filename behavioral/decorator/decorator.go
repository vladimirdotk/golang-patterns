package decorator

//Beverage - drink interface
type Beverage interface {
	GetDescription() string
	GetCost() float32
}

//Espresso coffee
type Espresso struct{}

//GetDescription - Espresso description
func (es *Espresso) GetDescription() string {
	return "Espresso"
}

//GetCost - Espresso cost
func (es *Espresso) GetCost() float32 {
	return 1.99
}

//Mocha addition
type Mocha struct {
	beverage    Beverage
	description string
	cost        float32
}

//GetDescription - Mocha description
func (m *Mocha) GetDescription() string {
	return m.beverage.GetDescription() + ", " + m.description
}

//GetCost - Mocha cost
func (m *Mocha) GetCost() float32 {
	return m.beverage.GetCost() + m.cost
}
