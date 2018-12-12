package mahjong

import (
	"strconv"
)

// NewTile creates a new tile
func NewTile(suit int, value uint) Tile {
	return Tile {Suit: suit, Value: value}
}

// Tile represents a mahjong tile
type Tile struct {
	Suit  int
	Value uint
}

var suitStr = []string{"c", "d", "b"}
var suitMap = map[string]int {
	"c": 0,
	"d": 1,
	"b": 2,
}

// ToString converts tile to string
func (tile Tile) ToString() string {
	return IF(tile.Suit < 0 || tile.Suit >= 3, "", suitStr[tile.Suit] + strconv.Itoa(int(tile.Value + 1))).(string)
}

// StringArrayToTileArray converts string array to tile array
func StringArrayToTileArray(tiles []string) []Tile {
	var res []Tile
	for _, tile := range tiles {
		res = append(res, StringToTile(tile))
	}
	return res
}

// StringToTile converts string to tile
func StringToTile(tile string) Tile {
	if tile == "" {
		return NewTile(-1, 0)
	}
	r    := []rune(tile)
	s    := string(r[0])
	v, _ := strconv.Atoi(string(r[1]))
	flag := false

	for _, suit := range suitStr {
		if s == suit {
			flag = true
			break
		}
	}
	return IF(!flag || v < 1, NewTile(-1, 0), NewTile(suitMap[s], uint(v - 1))).(Tile)
}

// IsValidTile checks if tile string is vaild
func IsValidTile(tile string) bool {
	if tile == "" {
		return false
	}
	r    := []rune(tile)
	s    := string(r[0])
	v, _ := strconv.Atoi(string(r[1]))
	flag := false
	for _, suit := range suitStr {
		if s == suit {
			flag = true
			break
		}
	}
	return flag && v > 0 && v <= 9
}