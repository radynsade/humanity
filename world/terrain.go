package world

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

type Terrain struct {
	landscape Landscape
	biome     Biome
}

func NewTerrain(landscape Landscape, biome Biome) *Terrain {
	return &Terrain{landscape, biome}
}

func (t *Terrain) Landscape() Landscape {
	return t.landscape
}

func (t *Terrain) Biome() Biome {
	return t.biome
}
