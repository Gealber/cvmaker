package layout

import "fmt"

// Layout of the resume.
type Layout struct {
	Regions []Region
}

// New creates a new layout with the specified regions.
// Regions shouldn't collide, otherwise you will receive an error.
func New(regions []Region) (*Layout, error) {
	for i := 0; i < len(regions)-1; i++ {
		for j := i + 1; j < len(regions); j++ {
			if regions[i].Collides(regions[j]) {
				return nil, errCollision(regions[i], regions[j])
			}
		}
	}

	return &Layout{Regions: regions}, nil
}

func errCollision(r1, r2 Region) error {
	return fmt.Errorf(
		"%s:\n %+v\n %+v\n",
		errRegionsCollidesMsg, r1, r2)
}

// DefaultLayout returns a default layout.
func DefaultLayout() (*Layout, error) {
	return New(defaultRegions)
}
