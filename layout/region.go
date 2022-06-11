package layout

import "github.com/jung-kurt/gofpdf"

// point in the region.
type point struct {
	X float64
	Y float64
}

// Region in the layout.
type Region struct {
	Name        string
	TopLeft     point
	BottomRight point
}

type Polygon []gofpdf.PointType

// New initiates a new region.
func newRegion(topCorner, bottomCorner point, name string) Region {
	return Region{
		TopLeft:     topCorner,
		BottomRight: bottomCorner,
		Name:        name,
	}
}

// Belongs checks if a point is inse the region.
func (r Region) Belongs(p point) bool {
	return pointInRegion(p, r)
}

// Collides checks if two regions collides, region r with r2.
func (r Region) Collides(r2 Region) bool {
	return collides(r, r2)
}

// Polygon returns a list of PointType that represents the region.
func (r Region) Polygon() Polygon {
	polygon := []gofpdf.PointType{
		{
			X: r.TopLeft.X,
			Y: r.TopLeft.Y,
		},
		{
			X: r.BottomRight.X,
			Y: r.TopLeft.Y,
		},
		{
			X: r.BottomRight.X,
			Y: r.BottomRight.Y,
		},
		{
			X: r.TopLeft.X,
			Y: r.BottomRight.Y,
		},
	}

	return polygon
}

func (r Region) Width() float64 {
	return r.BottomRight.X - r.TopLeft.X
}

func (r Region) Height() float64 {
	return r.BottomRight.Y - r.TopLeft.Y
}
