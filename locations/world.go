package locations

type WorldMap struct {
	Cities    []City
	MaxPlusX  int
	MaxPlusY  int
	MaxMinusX int
	MaxMinusY int
	Age       int
}

func BuildWorldMap() WorldMap {
	return WorldMap{
		Cities:    GenerateCities(),
		MaxPlusX:  100,
		MaxPlusY:  100,
		MaxMinusX: -100,
		MaxMinusY: -100,
		Age:       0,
	}
}

func (w WorldMap) Tick() WorldMap {
	for i, city := range w.Cities {
		c := city.Tick()
		w.Cities[i] = c
	}
	w.Age += 1
	return w
}
