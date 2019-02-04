package items

import (
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

func (m Market) Tick() Market {
	for i, commodity := range m.Commodities {
		c := commodity.Tick()
		m.Commodities[i] = c
	}
	return m
}
