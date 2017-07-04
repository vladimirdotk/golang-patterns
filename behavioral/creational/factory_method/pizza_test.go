package pizza

import "testing"

func TestPizza(t *testing.T) {
	stores := [2]string{"Chicago", "New York"}
	for _, store := range stores {
		if _, ok := getPizzaFromStore(store); !ok {
			t.Errorf("Fail to get pizza from " + store + " store")
		}
	}
}
