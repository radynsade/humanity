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

type Landscape uint8

const (
	LandscapeOcean = Landscape(iota)
	LandscapeSea
	LandscapePlain
	LandscapeHill
	LandscapeMountain
)

type Biome uint8

const (
	GrassBiome = Biome(iota)
	SavannahBiome
	DesertBiome
	TundraBiome
	IceBiome
)

type Cell struct {
	grid      *Grid
	x         uint8
	y         uint8
	landscape Landscape
	biome     Biome
}

func newCell(grid *Grid, x, y uint8) Cell {
	return Cell{
		grid: grid,
		x:    x,
		y:    y,
	}
}

func byteDiff(a, b uint8) uint8 {
	if a > b {
		return a - b
	}

	return b - a
}

func (c *Cell) X() uint8 {
	return c.x
}

func (c *Cell) Y() uint8 {
	return c.y
}

func (c *Cell) Landscape() Landscape {
	return c.landscape
}

func (c *Cell) SetLandscape(landscape Landscape) {
	c.landscape = landscape
}

func (c *Cell) Biome() Biome {
	return c.biome
}

func (c *Cell) SetBiome(biome Biome) {
	c.biome = biome
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

func (c *Cell) Distance(cell *Cell) uint8 {
	dy := byteDiff(c.y, cell.y)
	dx := byteDiff(c.x, cell.x)

	// Even-r specific adjustment.
	if (c.y%2 == 0 && cell.y%2 == 1 && c.x < cell.x) ||
		(c.y%2 == 1 && cell.y%2 == 0 && c.x > cell.x) {
		return max(dy, dx+dy/2)
	}

	return max(dx, dy+(dx+1)/2)
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
