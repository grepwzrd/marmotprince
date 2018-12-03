package locations

type WorldMap struct {
	Cities    []City
	MaxPlusX  int
	MaxPlusY  int
	MaxMinusX int
	MaxMinusY int
}

func BuildWorldMap() WorldMap {
	return WorldMap{
		Cities:    GenerateCities(),
		MaxPlusX:  100,
		MaxPlusY:  100,
		MaxMinusX: -100,
		MaxMinusY: -100,
	}
}
