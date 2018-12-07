package locations

import (
	"github.com/grepwzrd/marmotprince/items"
)

func (c City) Coordinates() [2]int {
	return [2]int{c.PositionX, c.PositionY}
}

type City struct {
	Name          string
	Market        items.Market
	PositionX     int
	PositionY     int
	PrimaryExport items.Commodity
	PrimaryImport items.Commodity
}

func buildCity(name string, positionX int, positionY int) City {
	market := items.GenerateMarket()
	return City{
		Name:          name,
		PositionX:     positionX,
		PositionY:     positionY,
		Market:        market,
		PrimaryExport: market.PrimaryExport,
		PrimaryImport: market.PrimaryImport,
	}
}

func GenerateCities() []City {
	cities := []City{}
	marmoa := buildCity("Marmoa", 4, 20)
	olympus := buildCity("Olympus", -54, 84)
	farburrow := buildCity("Farburrow", -70, -57)
	cities = append(cities, marmoa)
	cities = append(cities, olympus)
	cities = append(cities, farburrow)
	return cities
}

func (c City) Tick() {
	c.Market.Tick()
}
