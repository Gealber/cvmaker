package maker

const (
	defaultMargin       = 10
	defaultRegionMargin = 1
	smallFontSize       = 10
	defaultStep         = 10
	smalltStep          = 5
	defaultCellWidth    = 40
	defaultCellHeight   = 40

	defaultImageSize      = 30
	defaultHeaderFontSize = 30

	defaultFontSize  = 14
	mediumFontSize   = 16
	largeFontSize    = 18
	defaultLineWidth = 0.4
)

type colorRGB [3]int

var (
	// colors.
	grey    colorRGB = [3]int{155, 155, 155}
	black   colorRGB = [3]int{0, 0, 0}
	redSoft colorRGB = [3]int{241, 6, 6}
)
