package layout

const (
	unit                  = 10
	errRegionsCollidesMsg = "regions colliding"
)

var (
	defaultRegions = []Region{
		// Basics section.
		newRegion(point{unit, unit}, point{15 * unit, 5 * unit}, "Header"),
		newRegion(point{16 * unit, unit}, point{20 * unit, 5 * unit}, "Image"),
		newRegion(point{unit, 5.5 * unit}, point{20 * unit, 9 * unit}, "Summary"),
		newRegion(point{unit, 9.1 * unit}, point{20 * unit, 14 * unit}, "Social Networks"),
		// Education section.
		newRegion(point{unit, 14.3 * unit}, point{20 * unit, 17 * unit}, "Education"),
		// Work section.
		newRegion(point{unit, 17.3 * unit}, point{20 * unit, 27 * unit}, "Work"),
	}
)
