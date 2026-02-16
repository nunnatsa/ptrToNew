package pointer

import "k8s.io/utils/pointer"

var (
	x = pointer.Int32(42) // want `replace pointer\.Int32\(42\) with the "new\(\)" built-in function: new\(int32\(42\)\)`
	_ = pointer.Int32Deref(x, 42)
)
