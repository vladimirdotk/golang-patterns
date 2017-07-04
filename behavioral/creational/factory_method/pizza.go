package pizza

type pizza struct {
	name   string
	cheese string
	taste  string
}

type pizzaStore interface {
	deliver() string
	makePizza() pizza
	getSmell() string
}

type newYorkPizzaStore struct {
	pizzaProduct pizza
}

func (nyps *newYorkPizzaStore) makePizza() pizza {
	newYorkPizza := pizza{}
	newYorkPizza.name = "New York Pizza"
	newYorkPizza.cheese = "Gouda"
	newYorkPizza.taste = "Sweet"
	nyps.pizzaProduct = newYorkPizza
	return newYorkPizza
}

func (nyps *newYorkPizzaStore) deliver() string {
	return "Delivering " + nyps.pizzaProduct.name
}

func (nyps *newYorkPizzaStore) getSmell() string {
	return "Smells like " + nyps.pizzaProduct.cheese + " and " + nyps.pizzaProduct.taste
}

type chicagoPizzaStore struct {
	pizzaProduct pizza
}

func (cps *chicagoPizzaStore) makePizza() pizza {
	chicagoPizza := pizza{}
	chicagoPizza.name = "Chicago Pizza"
	chicagoPizza.cheese = "Mozzarella"
	chicagoPizza.taste = "Bitter"
	cps.pizzaProduct = chicagoPizza
	return chicagoPizza
}

func (cps *chicagoPizzaStore) deliver() string {
	return "Delivering " + cps.pizzaProduct.name
}

func (cps *chicagoPizzaStore) getTaste() string {
	return "Smells like " + cps.pizzaProduct.cheese + " and " + cps.pizzaProduct.taste
}

func getPizzaFromStore(pizzaType string) (pizza, bool) {
	if pizzaType == "New York" {
		return new(newYorkPizzaStore).makePizza(), true
	} else if pizzaType == "Chicago" {
		return new(chicagoPizzaStore).makePizza(), true
	}
	return pizza{}, false
}
