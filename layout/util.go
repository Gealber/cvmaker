package layout

func collides(r1, r2 region) bool {
	// r1 tottally on the left side of r2.
	leftSide := r1.TopLeft.X > r2.BottomRight.X
	// r1 tottally on the right side of r2.
	rightSide := r1.BottomRight.X < r2.TopLeft.X
	// r1 tottally bellow of r2.
	bellow := r1.TopLeft.Y > r2.BottomRight.Y
	// r1 tottally above of r2.
	above := r1.BottomRight.Y < r2.TopLeft.Y

	if leftSide || rightSide || bellow || above {
		return false
	}

	return true
}

func pointInRegion(p point, r region) bool {
	bellowTopLeft := p.X > r.TopLeft.X && p.Y > r.TopLeft.Y
	aboveBottomRight := p.X < r.BottomRight.X && p.Y < r.BottomRight.Y

	return bellowTopLeft && aboveBottomRight
}
