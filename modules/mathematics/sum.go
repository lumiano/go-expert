package mathematic

type Summable interface {
	~int | ~float64
}

type Calculator struct {
	Accumulator int
	//	hidden      bool
}

func Sum[T Summable](a, b T) T {

	return a + b
}

var HelloWorld = "Hello World"

const IsVisible = true
