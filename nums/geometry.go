package nums

import "image"

// AspectRatio returns geometry that conforms to the aspect ratio.
func AspectRatio(srcRect image.Point, toResize uint64) image.Point {
	w, h := int(toResize), getRatioSize(int(toResize), srcRect.Y, srcRect.X)
	if srcRect.X < srcRect.Y {
		w, h = getRatioSize(int(toResize), srcRect.X, srcRect.Y), int(toResize)
	}
	return image.Point{w, h}
}

func getRatioSize(a, b, c int) int {
	d := a * b / c
	return (d + 1) & -1
}
