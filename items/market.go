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
