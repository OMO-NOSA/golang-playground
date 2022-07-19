package structinterface

type Rectangle struct {
	width float64
	height float64
}

func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

func Area(width float64, height float64) float64 {
	return width * height
}




//struct is a named collection of fields where data can be stored.