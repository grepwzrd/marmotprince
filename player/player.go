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
		Positionx: 4, // randomize starting loc from diff town locs
		Positiony: 20,
		Gold:      420, // blaze it
	}
}
