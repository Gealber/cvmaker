package layout

const (
	unit                  = 10
	errRegionsCollidesMsg = "regions colliding"
)

var (
	defaultRegions = []region{
		newRegion(point{unit, unit}, point{15 * unit, 5 * unit}, "Basics"),
		newRegion(point{16 * unit, unit}, point{20 * unit, 5 * unit}, "Image"),
		newRegion(point{unit, 6 * unit}, point{20 * unit, 10 * unit}, "Basics"),
		newRegion(point{unit, 11 * unit}, point{20 * unit, 15 * unit}, "Social networks"),
		newRegion(point{unit, 16 * unit}, point{20 * unit, 20 * unit}, "Education"),
		newRegion(point{unit, 21 * unit}, point{20 * unit, 28 * unit}, "Work experience"),
		newRegion(point{unit, 29 * unit}, point{20 * unit, 29 * unit}, "Footer Line"),
	}
)
