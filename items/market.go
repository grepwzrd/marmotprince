package items

import (
	"fmt"
	"math/rand"
)

type Market struct {
	Commodities   []Commodity
	PrimaryExport Commodity
	PrimaryImport Commodity
}

func GenerateMarket() Market {
	commodities := BuildCommoditiesList()
	return Market{
		Commodities:   commodities,
		PrimaryExport: commodities[rand.Intn(len(commodities))],
		PrimaryImport: commodities[rand.Intn(len(commodities))],
	}
}

func plusHundo(c *Commodity) {
	c.price = c.price + 400
}

func (m Market) Tick() {
	for _, c := range m.Commodities {
		fmt.Println(c.Name)
		plusHundo(&c)
		fmt.Println(c.price)
		//FluctuatePrice(c)
	}
}
