package worldmap

import "math"

type Cell struct {
	worldMap *WorldMap
}

func (c *Cell) Left() {

}

func (c *Cell) Right() {

}

func (c *Cell) TopLeft() {

}

func (c *Cell) TopRight() {

}

func (c *Cell) BottomLeft() {

}

func (c *Cell) BottomRight() {

}

type WorldMap struct {
	cells  [math.MaxUint16]*Cell
	width  uint8
	height uint8
}

func NewWorldMap(width, height uint8) *WorldMap {
	var (
		i, j  uint8
		cells [math.MaxUint16]*Cell
	)

	worldMap := &WorldMap{}

	for i = range width {
		for j = range height {
			cells[i+j] = &Cell{worldMap}
		}
	}

	worldMap.cells = cells
	worldMap.width = width
	worldMap.height = height
	return worldMap
}

func (wm WorldMap) Cell(x, y uint8) *Cell {

	return wm.cells[x*y]
}
