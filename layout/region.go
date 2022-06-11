package layout

// point in the region.
type point struct {
	X float64
	Y float64
}

// region in the layout.
type region struct {
	Name        string
	TopLeft     point
	BottomRight point
}

// New initiates a new region.
func newRegion(topCorner, bottomCorner point, name string) region {
	return region{
		TopLeft:     topCorner,
		BottomRight: bottomCorner,
		Name:        name,
	}
}

// Belongs checks if a point is inse the region.
func (r region) Belongs(p point) bool {
	return pointInRegion(p, r)
}

// Collides checks if two regions collides, region r with r2.
func (r region) Collides(r2 region) bool {
	return collides(r, r2)
}
