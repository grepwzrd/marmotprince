package player

type Merchant struct {
	Name      string
	Positionx int
	Positiony int
	Gold      int
}

func InitializePlayer() Merchant {
	return Merchant{
		Name:      GenerateName(),
		Positionx: 0, // randomize starting loc from diff town locs
		Positiony: 0,
		Gold:      420, // blaze it
	}
}
