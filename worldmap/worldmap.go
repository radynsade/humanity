package worldmap

import "math"

type Cell struct {
	worldMap *WorldMap
	x        uint8
	y        uint8
}

func (c *Cell) X() uint8 {
	return c.x
}

func (c *Cell) Y() uint8 {
	return c.y
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
		x, y  uint8
		cells [math.MaxUint16]*Cell
	)

	worldMap := &WorldMap{}

	for x = range width {
		for y = range height {
			// "x" and "y" are starting here from 0, so we don't need to - 1.
			cells[y*worldMap.width+x] = &Cell{
				worldMap: worldMap,
				x:        x + 1,
				y:        y + 1,
			}
		}
	}

	worldMap.cells = cells
	worldMap.width = width
	worldMap.height = height
	return worldMap
}

func (wm *WorldMap) Cell(x, y uint8) *Cell {
	return wm.cells[(y-1)*wm.width+x]
}
