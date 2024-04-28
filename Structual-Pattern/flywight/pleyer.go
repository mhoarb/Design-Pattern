package main

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := getDressFactorysingleInstance().getDressByType(dressType)
	return &Player{
		dress:      dress,
		playerType: playerType,
	}
}
