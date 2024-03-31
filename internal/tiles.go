package internal

import (
	"crypto/rand"
	"math/big"
	"slices"
)

var tiles = []string{
	"b1", "b2", "b3", "b4", "b5", "b6", "b7", "b8", "b9",
	"n1", "n2", "n3", "n4", "n5", "n6", "n7", "n8", "n9",
	"p1", "p2", "p3", "p4", "p5", "p6", "p7", "p8", "p9",
	"n", "w", "s", "e", "red", "green", "white",
}

var allTiles = slices.Concat(tiles, tiles, tiles, tiles)

type Player struct {
	Points    int
	Hand      []string
	Discard   []string
	Pon       []string
	ClosedKan []string
	OpenKan   []string
	Chi       []string
	InRiichi  bool
}

type Game struct {
	GameTiles         []string
	Players           [4]Player
	TurnIndex         int
	TileIndex         int
	WallTiles         []string
	WallRevealedTiles int
}

func (g *Game) TakeTiles(n int) []string {
	res := g.GameTiles[g.TileIndex : g.TileIndex+n]
	g.TileIndex += n
	return res
}

func NewPlayer(handTiles []string) Player {
	return Player{
		Points:    25_000,
		Hand:      handTiles,
		Discard:   []string{},
		Pon:       []string{},
		ClosedKan: []string{},
		OpenKan:   []string{},
		Chi:       []string{},
		InRiichi:  false,
	}
}

func PanicOnErrR[T any](res T, err error) T {
	if err != nil {
		panic(err)
	}
	return res
}

func shuffle(slice []string) {
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		j := PanicOnErrR(rand.Int(rand.Reader, big.NewInt(int64(i+1))))
		jVal := int(j.Int64())
		slice[i], slice[jVal] = slice[jVal], slice[i]
	}
}

func NewGame() Game {
	gameTiles := make([]string, len(allTiles))
	copy(gameTiles, allTiles)
	shuffle(gameTiles)
	game := Game{
		GameTiles:         gameTiles,
		Players:           [4]Player{},
		TurnIndex:         0,
		TileIndex:         0,
		WallTiles:         nil,
		WallRevealedTiles: 1,
	}
	game.WallTiles = game.TakeTiles(10)
	for i := range len(game.Players) {
		game.Players[i].Hand = game.TakeTiles(13)
	}
	return game
}
