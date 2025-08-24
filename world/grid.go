// Pointy-top hexagonal even-r grid.
package world

import "math"

type Direction uint8

const (
	DirectionLeft = Direction(iota)
	DirectionRight
	DirectionTopLeft
	DirectionTopRight
	DirectionBottomLeft
	DirectionBottomRight
)

type Layer uint8

const (
	LayerTerrain = Layer(iota)
	LayerFlora
	LayerResource
	LayerBuilding
	LayerUnit
	LayerWeather
)

type Object struct {
	cell  *Cell
	layer Layer
}

type Cell struct {
	grid    *Grid
	x       uint8
	y       uint8
	objects map[Layer][math.MaxUint8]Object
}

func newCell(grid *Grid, x, y uint8) Cell {
	return Cell{
		grid:    grid,
		x:       x,
		y:       y,
		objects: map[Layer][math.MaxUint8]Object{},
	}
}

func (c *Cell) X() uint8 {
	return c.x
}

func (c *Cell) Y() uint8 {
	return c.y
}

func (c *Cell) Left() *Cell {
	if c.x == 0 {
		return nil
	}

	return c.grid.Cell(c.x-1, c.y)
}

func (c *Cell) Right() *Cell {
	if c.x == c.grid.width-1 {
		return nil
	}

	return c.grid.Cell(c.x+1, c.y)
}

func (c *Cell) TopLeft() *Cell {
	if c.x == 0 || c.y == 0 {
		return nil
	}

	return c.grid.Cell(c.x-1, c.y-1)
}

func (c *Cell) TopRight() *Cell {
	if c.x == c.grid.width-1 || c.y == 0 {
		return nil
	}

	return c.grid.Cell(c.x, c.y-1)
}

func (c *Cell) BottomLeft() *Cell {
	if c.x == 0 || c.y == c.grid.height-1 {
		return nil
	}

	return c.grid.Cell(c.x-1, c.y+1)
}

func (c *Cell) BottomRight() *Cell {
	if c.x == c.grid.width-1 || c.y == c.grid.height-1 {
		return nil
	}

	return c.grid.Cell(c.x, c.y+1)
}

func (c *Cell) Distance(cell *Cell) uint16 {
	return c.x
}

type Grid struct {
	cells  [math.MaxUint16]Cell
	width  uint8
	height uint8
}

func NewGrid(width, height uint8) *Grid {
	var (
		x, y  uint8
		cells [math.MaxUint16]Cell
	)

	grid := &Grid{}

	for x = range width {
		for y = range height {
			cells[y*grid.width+x] = newCell(grid, x, y)
		}
	}

	grid.cells = cells
	grid.width = width
	grid.height = height
	return grid
}

func (g *Grid) Cell(x, y uint8) *Cell {
	return &g.cells[y*g.width+x]
}
