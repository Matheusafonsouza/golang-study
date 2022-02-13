package main

type Retangle struct {
	height float64
	width  float64
}

func perimeter(retangle Retangle) float64 {
	return 2 * (retangle.width + retangle.height)
}

func area(retangle Retangle) float64 {
	return retangle.width * retangle.height
}
