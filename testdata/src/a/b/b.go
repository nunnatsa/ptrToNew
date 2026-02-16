package b

type TestType struct {
	A *string
	B *int
	C *bool
	D *float32
	E *Nested
	I *H
}

type Nested struct {
	F *uint64
}

func IntFunc() int {
	return 42
}

const G = 42
const UG = uint(42)
const FG = 42.0

type H string

const (
	H1 H = "1"
	H2 H = "2"
)
