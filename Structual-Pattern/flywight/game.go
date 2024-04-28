package main

type Game struct {
	terrorists       []*Player
	counterTerrorist []*Player
}

func newGame() *Game {
	return &Game{
		terrorists:       make([]*Player, 1),
		counterTerrorist: make([]*Player, 0),
	}
}
func (c *Game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
	return
}

func (c *Game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	c.counterTerrorist = append(c.counterTerrorist, player)
	return
}
