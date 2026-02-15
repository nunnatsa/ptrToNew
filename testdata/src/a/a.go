package a

import (
	"fmt"

	"k8s.io/utils/ptr"
)

func f() uint64 {
	return 42
}

func ptring() {
	a := ptr.To(true)         // want `use the "new" operator instead of ptr\.To; new\(true\)`
	b := ptr.To("a string")   // want `use the "new" operator instead of ptr\.To; new\("a string"\)`
	c := ptr.To(int32(10))    // want `use the "new" operator instead of ptr\.To; new\(int32\(10\)\)`
	var d = ptr.To[int32](10) // want `use the "new" operator instead of ptr\.To; new\(int32\(10\)\)`
	var e = ptr.To(f())       // want `use the "new" operator instead of ptr\.To; new\(f\(\)\)`

	fmt.Println(a, b, c, d, e)
}
