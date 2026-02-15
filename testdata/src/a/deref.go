package a

import (
	"fmt"

	"k8s.io/utils/ptr"
)

func useDeref() {
	a := ptr.To(true) // want `use the "new" operator instead of ptr\.To; new\(true\)`
	if val := ptr.Deref(a, false); val {
		fmt.Println(val)
	}
}
