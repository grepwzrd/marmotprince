package items

import (
	"math/rand"
)

type Commodity struct {
	Name     string
	price    int
	UnitName string
	Quantity int
}

func generatePrice(baseRange []int) int {
	return baseRange[rand.Intn(len(baseRange))]
}

func generateQuantity() int {
	return rand.Intn(135)
}

func generateCommodity(name string, priceRange []int, unitName string) Commodity {
	return Commodity{
		Name:     name,
		price:    generatePrice(priceRange),
		UnitName: unitName,
		Quantity: generateQuantity(),
	}
}

func BuildCommoditiesList() []Commodity {
	c := []Commodity{}
	c = append(c, generateCommodity("popcorn", []int{1, 1, 1, 2}, "kernel"))
	c = append(c, generateCommodity("juniper berries", []int{2, 3, 4}, "bushel"))
	c = append(c, generateCommodity("funherbs", []int{10, 11, 12}, "gram"))
	c = append(c, generateCommodity("bark trinket", []int{15, 21, 22}, "each"))
	c = append(c, generateCommodity("pinecones", []int{6, 7}, "each"))
	c = append(c, generateCommodity("acorns", []int{1, 2, 3}, "cheekpouch"))
	return c
}

func (c Commodity) Price() int {
	return c.price
}

func FluctuatePrice(c Commodity) {
	if rand.Intn(3) == 0 {
		updatePrice(&c, c.price+1)
	} else if rand.Intn(4) == 0 {
		updatePrice(&c, c.price-1)
	}
}

func updatePrice(c *Commodity, newPrice int) {
	a := Commodity{Name: c.Name, UnitName: c.UnitName, Quantity: c.Quantity, price: newPrice}
	*c = a
}
